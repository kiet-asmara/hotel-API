package model

type Booking struct {
	Booking_id    int     `json:"booking_id" gorm:"primary_key"`
	User_id       int     `json:"user_id"`
	Room_id       int     `json:"room_id"`
	Checkin_date  string  `json:"checkin_date"`
	Checkout_date string  `json:"checkout_date"`
	Total_price   float32 `json:"total_price"`
	Paid          bool    `json:"paid"`
}
