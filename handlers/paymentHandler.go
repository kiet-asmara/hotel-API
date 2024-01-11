package handlers

import (
	"hotel/helpers"
	"hotel/service"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DepositHandler(c echo.Context) error {
	userID, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}
	var input service.DepositInput

	err = c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	invoice, err := h.Service.Deposit(input.Deposit_amount, userID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "deposit invoice created",
		"invoice": invoice,
	})
}

func (h *Handler) DepositRefreshHandler(c echo.Context) error {
	userID, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	depositHistory, totalDeposits, err := h.Service.DepositRefresh(userID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"deposit_amount": totalDeposits,
		"deposits":       depositHistory,
	})
}
