package model

type Room struct {
	Room_id      int        `json:"room_id" gorm:"primary_key"`
	Room_type_id int        `json:"room_type_id"`
	Status       bool       `json:"available"`
	Room_type    *Room_type `json:"room_type,omitempty"`
}

type Room_type struct {
	Room_type_id    int     `json:"room_type_id" gorm:"primary_key"`
	Room_name       string  `json:"room_name"`
	Description     string  `json:"description"`
	Price_per_night float32 `json:"price_per_night"`
}
