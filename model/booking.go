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

// func (booking Booking) RemoveTimezoneBooking() (Booking, error) {
// 	// parse string w/timezone to YYYY-MM-DD
// 	checkin, err := time.Parse("2006-01-02", booking.Checkin_date)
// 	if err != nil {
// 		return Booking{}, err
// 	}
// 	fmt.Println(booking)
// 	fmt.Println(checkin)
// 	fmt.Println(checkin.Format("2006-01-02"))

// 	checkout, err := time.Parse("2006-01-02", booking.Checkout_date)
// 	if err != nil {
// 		return Booking{}, err
// 	}

// 	// populate data
// 	res := Booking{
// 		Booking_id:    booking.Booking_id,
// 		User_id:       booking.User_id,
// 		Room_id:       booking.Room_id,
// 		Checkin_date:  checkin.Format("2006-01-02"),
// 		Checkout_date: checkout.Format("2006-01-02"),
// 		Total_price:   booking.Total_price,
// 		Paid:          booking.Paid,
// 	}
// 	return res, nil
// }
