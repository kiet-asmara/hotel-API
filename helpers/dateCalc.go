package helpers

import (
	"time"
)

func DatesBetween(earlyDate, laterDate string) (float32, error) {
	early, err := time.Parse("2006-01-02", earlyDate)
	if err != nil {
		return 0, err
	}

	later, err := time.Parse("2006-01-02", laterDate)
	if err != nil {
		return 0, err
	}

	Duration := later.Sub(early).Hours() / 24
	return float32(Duration), nil
}

// func RemoveTimezoneBooking(booking model.Booking) (model.Booking, error) {
// 	// parse string w/timezone to YYYY-MM-DD
// 	checkin, err := time.Parse("2006-01-02", booking.Checkin_date)
// 	if err != nil {
// 		return model.Booking{}, err
// 	}
// 	fmt.Println(booking)
// 	fmt.Println(checkin)
// 	fmt.Println(checkin.Format("2006-01-02"))

// 	checkout, err := time.Parse("2006-01-02", booking.Checkout_date)
// 	if err != nil {
// 		return model.Booking{}, err
// 	}

// 	// populate data
// 	res := model.Booking{
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
