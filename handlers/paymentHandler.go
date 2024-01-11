package handlers

import (
	"hotel/helpers"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DepositHandler(c echo.Context) error {
	userID, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	return c.JSON(http.StatusOK)
}
