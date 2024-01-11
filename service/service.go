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
	Deposit_amount float32 `json:"deposit_amount" validate:"required"`
}

// xendit objects
type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type XenditRequest struct {
	External_id      string   `json:"external_id"`
	Amount           float32  `json:"amount"`
	Description      string   `json:"description"`
	Invoice_duration int      `json:"invoice_duration"`
	Customer         Customer `json:"customer"`
	Items            []Item   `json:"items"`
}

type XenditResponse struct {
	Id          string         `json:"id"`
	External_id string         `json:"external_id"`
	User_id     string         `json:"user_id"`
	Status      string         `json:"status"`
	Amount      float32        `json:"amount"`
	Expiry_date string         `json:"expiry_date"`
	Invoice_url string         `json:"invoice_url"`
	Customer    map[string]any `json:"customer"`
	// Created     string         `json:"created"`
	// Updated     string         `json:"updated"`
}
