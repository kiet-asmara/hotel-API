package handlers

import (
	"hotel/model"
	"hotel/service"
)

type Handler struct {
	Service *service.Service
}

// response outputs for swagger
type ErrResponse struct {
	Message string `json:"message"`
}

type RegisterResponse struct {
	Message string           `json:"message"`
	User    service.UserInfo `json:"user"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type DepositResponse struct {
	Message string                 `json:"message"`
	Invoice service.XenditResponse `json:"invoice"`
}

type GetDepositResponse struct {
	Deposit_amount float32         `json:"deposit_amount"`
	Deposits       []model.Deposit `json:"deposits"`
}

type BookingResponse struct {
	Bookings []model.Booking `json:"bookings"`
}

type PaymentReq struct {
	Payment_method string `json:"payment_method"`
}

type PaymentResponse struct {
	Message string        `json:"message"`
	Payment model.Payment `json:"payment"`
}

type PaymentsResponse struct {
	Payments []model.Payment `json:"payments"`
}

type RoomTyResp struct {
	Room_types []model.Room_type `json:"room_types"`
}

type RoomResp struct {
	Rooms []model.Room `json:"rooms"`
}

type RoomResp2 struct {
	Message string     `json:"message"`
	Room    model.Room `json:"room"`
}

type RoomResp3 struct {
	Message   string          `json:"message"`
	Room_type model.Room_type `json:"room_type"`
}

type RoomBookReq struct {
	Room_id       int    `json:"room_id" validate:"required"`
	Checkin_date  string `json:"checkin_date" validate:"required"`
	Checkout_date string `json:"checkout_date" validate:"required"`
}

type RoomBookResp struct {
	Message string        `json:"message"`
	Booking model.Booking `json:"booking"`
}

type RoomTypeCreate struct {
	Room_name       string  `json:"room_name" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	Price_per_night float32 `json:"price_per_night" validate:"required"`
	Available_rooms int     `json:"available_rooms"`
}
