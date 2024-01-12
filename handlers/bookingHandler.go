package handlers

import (
	"hotel/helpers"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Get all Bookings
// @Description  Get all and refresh user deposit history
// @Tags         Booking
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object}  handlers.BookingResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /bookings [Get]
func (h *Handler) ShowBookingHandler(c echo.Context) error {
	// get logged in user id
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	// get bookings
	bookings, err := h.Service.ShowUserBookings(claims.UserID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"bookings": bookings,
	})
}
