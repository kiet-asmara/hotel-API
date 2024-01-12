package handlers

import (
	"hotel/helpers"
	"hotel/service"
	"hotel/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Deposit
// @Description  Deposit to balance with Xendit invoice
// @Tags         User
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 data body service.DepositInput true "The input deposit struct"
// @Success      201  {object}  handlers.DepositResponse
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /users/deposit [Post]
func (h *Handler) DepositHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}
	var input service.DepositInput

	err = c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	invoice, err := h.Service.Deposit(input.Deposit_amount, claims.UserID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "deposit invoice created",
		"invoice": invoice,
	})
}

// @Summary      Get all Deposits
// @Description  Get all and refresh user deposit history
// @Tags         User
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object}  handlers.GetDepositResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /users/deposit [Get]
func (h *Handler) DepositRefreshHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	depositHistory, totalDeposits, err := h.Service.DepositRefresh(claims.UserID)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"deposit_amount": totalDeposits,
		"deposits":       depositHistory,
	})
}
