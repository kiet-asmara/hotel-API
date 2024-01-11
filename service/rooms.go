package service

import (
	"hotel/helpers"
	"hotel/model"
	"hotel/utils"
)

func (s *Service) GetRoomTypes() ([]model.Room_type, error) {
	roomTypes := []model.Room_type{}

	// tambah available rooms count?
	err := s.DB.Model(model.Room_type{}).Find(&roomTypes).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	return roomTypes, nil
}

func (s *Service) GetAvailableRooms(typeID string) ([]model.Room, error) {
	rooms := []model.Room{}

	err := s.DB.Model(model.Room{}).Where("room_type_id = ? AND status = true", typeID).Find(&rooms).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	return rooms, nil
}

func (s *Service) BookRoom(input BookingInput) (model.Booking, error) {
	// get room info
	room := model.Room{}
	err := s.DB.Model(model.Room{}).Preload("Room_type").First(&room, input.Room_id).Error
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	// calculate total price
	days, err := helpers.DatesBetween(input.Checkin_date, input.Checkout_date)
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrBadRequest, err) // bad request because date is not correct format
	}

	totalPrice := days * room.Room_type.Price_per_night

	// insert booking
	booking := model.Booking{
		User_id:       input.User_id,
		Room_id:       input.Room_id,
		Checkin_date:  input.Checkin_date,
		Checkout_date: input.Checkout_date,
		Total_price:   totalPrice,
		Paid:          false,
	}

	err = s.DB.Omit("booking_id").Create(&booking).Error
	if err != nil {
		return model.Booking{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return booking, nil
}
