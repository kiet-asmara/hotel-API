package service

import (
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
