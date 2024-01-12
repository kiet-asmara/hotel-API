package main

import (
	"hotel/config"
	"hotel/handlers"
	"hotel/service"
	"hotel/utils"

	_ "hotel/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Restful Hotel App API
// @version 1.0
// @description Hotel management system for users and admin.

// @contact.name Kiet Asmara
// @contact.url http://www.swagger.io/support
// @contact.email kiet123pascal@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	e := echo.New()

	// uncomment for local deployment
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		e.Logger.Fatal("failed db", err)
	}

	handler := &handlers.Handler{
		Service: &service.Service{DB: db},
	}

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:    1 << 10, // 1 KB
		LogErrorFunc: utils.LogError,
	}))

	user := e.Group("/users")
	user.POST("/register", handler.RegisterHandler)
	user.POST("/login", handler.LoginHandler)
	user.POST("/deposit", handler.DepositHandler, utils.AuthMiddleware)
	user.GET("/deposit", handler.DepositRefreshHandler, utils.AuthMiddleware)

	room := e.Group("/rooms")
	room.Use(utils.AuthMiddleware)
	room.GET("", handler.RoomTypeHandler)
	room.GET("/:id", handler.AvailableRoomHandler)
	room.POST("/book", handler.RoomBookingHandler)
	room.POST("/:id", handler.CreateRoomHandler)
	room.POST("/type", handler.CreateRoomTypeHandler)
	room.PUT("/type/:id", handler.UpdateRoomTypeHandler)
	room.DELETE("/:id", handler.DeleteRoomHandler)

	booking := e.Group("/bookings")
	booking.Use(utils.AuthMiddleware)
	booking.GET("", handler.ShowBookingHandler)
	booking.POST("/:id", handler.PayBookingHandler)
	booking.GET("/payments", handler.PaymentRefreshHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
