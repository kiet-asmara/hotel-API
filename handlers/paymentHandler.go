package handlers

import (
	"hotel/helpers"
	"hotel/service"
	"hotel/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary      Pay Booking By ID
// @Description  Pay for a booking by ID using deposit or Xendit
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Booking ID"
// @Param		 data body handlers.PaymentReq true "The input payment struct"
// @Success      201  {object}  handlers.PaymentResponse
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /bookings/:id [Post]
func (h *Handler) PayBookingHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}
	var input service.PaymentInput

	err = c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequest)
	}

	input.Booking_id = bookingID
	input.User_id = claims.UserID

	payment, err := h.Service.PayBooking(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "payment record created",
		"payment": payment,
	})
}

// @Summary      Get all payments
// @Description  Get all user payments
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object}  handlers.PaymentsResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /bookings/payments [Get]
func (h *Handler) PaymentRefreshHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	payments, err := h.Service.PaymentRefresh(claims.UserID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"payments": payments,
	})
}
