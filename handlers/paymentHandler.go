package handlers

import (
	"hotel/helpers"
	"hotel/service"
	"hotel/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

	return c.JSON(http.StatusOK, echo.Map{
		"message": "payment record created",
		"payment": payment,
	})
}

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
