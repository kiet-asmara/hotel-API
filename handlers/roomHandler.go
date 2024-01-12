package handlers

import (
	"hotel/helpers"
	"hotel/model"
	"hotel/service"
	"hotel/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) RoomTypeHandler(c echo.Context) error {
	roomTypes, err := h.Service.GetRoomTypes()
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"room_types": roomTypes,
	})
}

func (h *Handler) AvailableRoomHandler(c echo.Context) error {
	id := c.Param("id")

	rooms, err := h.Service.GetAvailableRooms(id)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"rooms": rooms,
	})
}

func (h *Handler) RoomBookingHandler(c echo.Context) error {
	// bind user input
	var input service.BookingInput
	err := c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrFailedBind)
	}

	// get logged in user id
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}
	input.User_id = claims.UserID

	// book room
	Booking, err := h.Service.BookRoom(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "booking created",
		"booking": Booking,
	})
}

func (h *Handler) CreateRoomHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	if claims.Role != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrUnauthorized)
	}

	// bind user input
	id := c.Param("id")

	room, err := h.Service.CreateRoom(id)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "room created",
		"room":    room,
	})
}

func (h *Handler) CreateRoomTypeHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	if claims.Role != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrUnauthorized)
	}

	// bind user input
	var input model.Room_type
	err = c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	roomType, err := h.Service.CreateRoomType(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "room type created",
		"room_type": roomType,
	})
}

func (h *Handler) UpdateRoomTypeHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	if claims.Role != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrUnauthorized)
	}

	// bind user input
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// bind user input
	var input model.Room_type
	err = c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	input.Room_type_id = id

	room, err := h.Service.UpdateRoomType(input)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "room type updated",
		"room_type": room,
	})
}

func (h *Handler) DeleteRoomHandler(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrInternalFailure)
	}

	if claims.Role != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrUnauthorized)
	}

	// bind user input
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = h.Service.DeleteRoom(id)
	if err != nil {
		apiErr := utils.FromError(err)
		return echo.NewHTTPError(apiErr.Status, apiErr.Message)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "room deleted",
	})
}
