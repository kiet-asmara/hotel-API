package service

import (
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

type RegisterInput struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Full_name string `json:"full_name" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Full_name      string  `json:"full_name"`
	Deposit_amount float32 `json:"deposit_amount"`
}

type RoomInfo struct {
	Room_id      int  `json:"room_id" gorm:"primary_key"`
	Room_type_id int  `json:"room_type_id"`
	Status       bool `json:"available"`
}

type BookingInput struct {
	User_id       int    `json:"user_id"`
	Room_id       int    `json:"room_id" validate:"required"`
	Checkin_date  string `json:"checkin_date" validate:"required"`
	Checkout_date string `json:"checkout_date" validate:"required"`
}

type DepositInput struct {
	Deposit_amount float32 `json:"deposit"`
}
