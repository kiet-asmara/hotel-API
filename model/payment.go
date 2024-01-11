package model

type Payment struct {
	Payment_id   int    `json:"payment_id"`
	Booking_id   int    `json:"booking_id"`
	Payment_date string `json:"payment_date"`
}

type Deposit struct {
	Deposit_id int     `json:"deposit_id"`
	User_id    int     `json:"user_id"`
	Amount     float32 `json:"amount"`
	Status     string  `json:"status"`
	Invoice_id string  `json:"invoice_id"`
	URL        string  `json:"url"`
}
