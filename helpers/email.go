package helpers

import (
	"fmt"
	"hotel/model"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(email, subject, content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "kiet@pascal.com")
	m.SetHeader("To", email)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(
		os.Getenv("SMTP_SERVER"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendSuccessPayment(email string, payment model.Payment) {
	subject := fmt.Sprintf("Payment info for booking id %d", payment.Booking_id)
	content := fmt.Sprintf(`<h1>Payment Info</h1>
				<p>Payment ID: %d</p>
				<p>Booking ID: %d</p>
				<p>Payment Date: %s</p>
				<p>Payment Method: %s</p>
				<p>Amount: %v</p>
				<p>Status: %s</p>
				<p>Invoice ID: %s
				<p>Invoice URL: %s`,
		payment.Payment_id,
		payment.Booking_id,
		payment.Payment_date,
		payment.Payment_method,
		payment.Amount,
		payment.Status,
		payment.Invoice_id,
		payment.URL)

	SendMail(
		email,
		subject,
		content,
	)
}
