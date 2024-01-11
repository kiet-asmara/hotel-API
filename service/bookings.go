package service

import (
	"fmt"
	"hotel/helpers"
	"hotel/model"
	"hotel/utils"
)

func (s *Service) ShowUserBookings(userID int) ([]model.Booking, error) {
	bookings := []model.Booking{}

	// tambah available rooms count?
	err := s.DB.Model(model.Booking{}).Find(&bookings).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	newBookings := []model.Booking{}
	for _, v := range bookings {
		fmt.Println(v)
		bookingNoTZ, err := helpers.RemoveTimezoneBooking(v) // error omitted because database always has correct format
		if err != nil {
			return nil, utils.NewError(utils.ErrInternalFailure, err)
		}
		newBookings = append(newBookings, bookingNoTZ)
	}

	return newBookings, nil
}
