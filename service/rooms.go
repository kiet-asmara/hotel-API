package service

import (
	"hotel/model"
	"hotel/utils"
	"strconv"
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

	err := s.DB.Model(model.Room{}).Where("room_type_id = ? AND available = true", typeID).Find(&rooms).Error
	if err != nil {
		return nil, utils.NewError(utils.ErrInternalFailure, err)
	}

	return rooms, nil
}

func (s *Service) CreateRoom(typeID string) (model.Room, error) {
	id, err := strconv.Atoi(typeID)
	if err != nil {
		return model.Room{}, utils.NewError(utils.ErrBadRequest, err)
	}
	room := model.Room{
		Room_type_id: id,
		Available:    true,
	}

	err = s.DB.Omit("Room_id").Create(&room).Error
	if err != nil {
		return model.Room{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return room, nil
}

func (s *Service) CreateRoomType(input model.Room_type) (model.Room_type, error) {

	if err := utils.Validate.Struct(input); err != nil {
		return model.Room_type{}, utils.NewError(utils.ErrFailedBind, err)
	}

	input.Available_rooms = 0

	err := s.DB.Omit("Room_type_id").Create(&input).Error
	if err != nil {
		return model.Room_type{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return input, nil
}

func (s *Service) UpdateRoomType(input model.Room_type) (model.Room_type, error) {

	if err := utils.Validate.Struct(input); err != nil {
		return model.Room_type{}, utils.NewError(utils.ErrFailedBind, err)
	}

	err := s.DB.Save(&input).Error
	if err != nil {
		return model.Room_type{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return input, nil
}

func (s *Service) DeleteRoom(roomID int) error {
	room := model.Room{
		Room_id: roomID,
	}

	err := s.DB.Delete(&room).Error
	if err != nil {
		return utils.NewError(utils.ErrInternalFailure, err)
	}

	return nil
}
