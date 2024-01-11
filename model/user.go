package model

type User struct {
	User_id        int     `json:"user_id" gorm:"primary_key"`
	User_type      int     `json:"user_type"`
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required"`
	Full_name      string  `json:"full_name" validate:"required"`
	Deposit_amount float32 `json:"deposit_amount"`
}
