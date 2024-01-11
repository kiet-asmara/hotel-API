package service

import (
	"fmt"
	"hotel/model"
	"hotel/utils"
	"strconv"
)

func (s *Service) Deposit(amount float32, userID int) (XenditResponse, error) {
	// make customer
	user := model.User{}

	err := s.DB.Model(model.User{}).First(&user, userID).Error
	if err != nil {
		return XenditResponse{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	id := strconv.Itoa(userID)

	// create request payload for invoice creation
	req := XenditRequest{
		External_id:      id,
		Amount:           amount,
		Description:      "Deposit invoice for " + user.Full_name,
		Invoice_duration: 86400,
		Customer: Customer{
			Name:  user.Full_name,
			Email: user.Email},
		Items: []Item{},
	}

	resp, err := createInvoice(req)
	if err != nil {
		return XenditResponse{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	// insert deposit info to DB
	deposit := model.Deposit{
		User_id:    userID,
		Amount:     resp.Amount,
		Status:     resp.Status,
		Invoice_id: resp.Id,
		URL:        resp.Invoice_url,
	}
	err = s.DB.Omit("Deposit_id").Create(&deposit).Error
	if err != nil {
		return XenditResponse{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	// return resp object
	return resp, nil
}

func (s *Service) DepositRefresh(userID int) ([]model.Deposit, float32, error) {
	// get all deposits
	deposits := []model.Deposit{}

	err := s.DB.Model(model.Deposit{}).Where("User_id = ?", userID).Find(&deposits).Error
	if err != nil {
		return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
	}
	if len(deposits) == 0 {
		return nil, 0, utils.NewError(utils.ErrNotFound, fmt.Errorf("you have no deposits"))
	}

	// get user info
	var user model.User
	err = s.DB.Model(model.User{}).First(&user, userID).Error
	if err != nil {
		return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
	}

	totalDeposit := user.Deposit_amount

	for _, v := range deposits {
		if v.Status == "PENDING" {
			resp, err := getInvoice(v.Invoice_id)
			if err != nil {
				return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
			}
			if resp.Status == "SETTLED" || resp.Status == "PAID" {
				// update deposit history
				err := s.DB.Model(model.Deposit{}).Where("deposit_id = ?", v.Deposit_id).Update("status", "PAID").Error
				if err != nil {
					return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
				}
				totalDeposit += v.Amount
			}
		}
	}

	if totalDeposit != user.Deposit_amount {
		err := s.DB.Model(&user).Update("deposit_amount", totalDeposit).Error
		if err != nil {
			return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
		}

		// query deposits again to see updated status
		err = s.DB.Model(model.Deposit{}).Where("User_id = ?", userID).Find(&deposits).Error
		if err != nil {
			return nil, 0, utils.NewError(utils.ErrInternalFailure, err)
		}
	}
	return deposits, totalDeposit, nil
}
