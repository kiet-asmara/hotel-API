package service

import (
	"errors"
	"fmt"
	"hotel/helpers"
	"hotel/model"
	"hotel/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func (s *Service) PayBooking(input PaymentInput) (model.Payment, error) {
	if err := utils.Validate.Struct(input); err != nil {
		return model.Payment{}, utils.NewError(utils.ErrFailedBind, err)
	}

	// get total price
	var booking model.Booking
	err := s.DB.Model(model.Booking{}).First(&booking, input.Booking_id).Error
	if err != nil {
		return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	// check if booking is paid
	if booking.Paid {
		return model.Payment{}, utils.NewError(utils.ErrBadRequest, fmt.Errorf("booking is already paid"))
	}

	// get user info
	var user model.User
	err = s.DB.Model(model.User{}).First(&user, input.User_id).Error
	if err != nil {
		return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	if input.Payment_method == "deposit" {

		// deduct from user deposit
		depositLeft := user.Deposit_amount - booking.Total_price
		if depositLeft < 0 {
			return model.Payment{}, utils.NewError(utils.ErrBadRequest, fmt.Errorf("insufficient balance"))
		}
		user.Deposit_amount = depositLeft

		// transaction begin
		tx := s.DB.Begin()

		err = tx.Model(&user).Update("deposit_amount", depositLeft).Error
		if err != nil {
			tx.Rollback()
			return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
		}

		// update booking status
		err = tx.Model(&booking).Update("paid", true).Error
		if err != nil {
			tx.Rollback()
			return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
		}

		// create payment data
		payment := model.Payment{
			Booking_id:     input.Booking_id,
			Payment_date:   time.Now().Format("2006-01-02"),
			Payment_method: input.Payment_method,
			Amount:         booking.Total_price,
			Status:         "PAID",
		}

		err = tx.Omit("Payment_id").Create(&payment).Error
		if err != nil {
			tx.Rollback()
			return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
		}

		tx.Commit()

		// send email
		helpers.SendSuccessPayment(user.Email, payment)

		return payment, nil
	}

	if input.Payment_method == "xendit" {
		// create invoice
		req := XenditRequest{
			External_id:      strconv.Itoa(input.Booking_id),
			Amount:           booking.Total_price,
			Description:      "Payment invoice for " + user.Full_name,
			Invoice_duration: 86400,
			Customer: Customer{
				Name:  user.Full_name,
				Email: user.Email},
			Items: []Item{}, // ga sempet query utk dapetin room type
		}
		resp, err := createInvoice(req)
		if err != nil {
			return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
		}

		// create payment data
		payment := model.Payment{
			Booking_id:     input.Booking_id,
			Payment_date:   "-infinity",
			Payment_method: input.Payment_method,
			Amount:         booking.Total_price,
			Status:         "PENDING",
			Invoice_id:     resp.Id,
			URL:            resp.Invoice_url,
		}

		err = s.DB.Omit("Payment_id").Create(&payment).Error
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) { // broken, need pgconn
				return model.Payment{}, utils.NewError(utils.ErrBadRequest, err)
			}
			return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
		}

		// send email
		helpers.SendSuccessPayment(user.Email, payment)

		return payment, nil

	}

	return model.Payment{}, utils.NewError(utils.ErrInternalFailure, err)
}

func (s *Service) PaymentRefresh(userID int) ([]model.Payment, error) {
	// get all payments
	payments := []model.Payment{}

	err := s.DB.Raw(`SELECT * FROM payments
		INNER JOIN bookings ON payments.booking_id = bookings.booking_id
		WHERE bookings.user_id = ?;`, userID).Scan(&payments).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}
	if len(payments) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, fmt.Errorf("you have no payments"))
	}

	// get user info
	var user model.User
	err = s.DB.Model(model.User{}).First(&user, userID).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	for _, v := range payments {
		if v.Status == "PENDING" {
			resp, err := getInvoice(v.Invoice_id)
			if err != nil {
				return nil, utils.NewError(utils.ErrInternalFailure, err)
			}
			if resp.Status == "SETTLED" || resp.Status == "PAID" {
				tx := s.DB.Begin()

				// update payment
				err := tx.Model(model.Payment{}).Where("payment_id = ?", v.Payment_id).Updates(model.Payment{Status: "PAID", Payment_date: time.Now().Format("2006-01-02")}).Error
				if err != nil {
					tx.Rollback()
					return nil, utils.NewError(utils.ErrInternalFailure, err)
				}

				// update booking status
				err = s.DB.Model(model.Booking{}).Where("booking_id = ?", v.Booking_id).Update("paid", true).Error
				if err != nil {
					tx.Rollback()
					return nil, utils.NewError(utils.ErrInternalFailure, err)
				}
				tx.Commit()

				// send email
				helpers.SendSuccessPayment(user.Email, v)
			}
		}
	}

	// query payments again to see updated status
	err = s.DB.Raw(`SELECT * FROM payments
		INNER JOIN bookings ON payments.booking_id = bookings.booking_id
		WHERE bookings.user_id = ?;`, userID).Scan(&payments).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	return payments, nil
}
