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

// @Summary      Get Room Types
// @Description  Get all room types
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object}  handlers.RoomTyResp
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms [Get]
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

// @Summary      Get Available Rooms By Type
// @Description  Get all available rooms by type id
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Room Type ID"
// @Success      200  {object}  handlers.RoomResp
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/:id [Get]
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

// @Summary      Book a room
// @Description  Book room by id
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 data body handlers.RoomBookReq true "The input booking struct"
// @Success      201  {object}  handlers.RoomBookResp
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/book [Post]
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

// @Summary      Create a room
// @Description  Create a room
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Room Type ID"
// @Success      201  {object}  handlers.RoomResp2
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/:id [Post]
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

// @Summary      Create a room type
// @Description  Create a room type
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 data body handlers.RoomTypeCreate true "The input room struct"
// @Success      201  {object}  handlers.RoomResp3
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/type [Post]
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

// @Summary      Update a room type
// @Description  Update a room type
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Room Type ID"
// @Param		 data body handlers.RoomTypeCreate true "The input room struct"
// @Success      201  {object}  handlers.RoomResp3
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/type/:id [Put]
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

// @Summary      Delete room
// @Description  Delete a room
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Room ID"
// @Success      201  {object}  handlers.ErrResponse
// @Failure      400  {object}  handlers.ErrResponse
// @Failure      401  {object}  handlers.ErrResponse
// @Failure      404  {object}  handlers.ErrResponse
// @Failure      500  {object}  handlers.ErrResponse
// @Router       /rooms/:id [Delete]
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
