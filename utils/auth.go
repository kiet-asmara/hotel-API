package utils

import (
	"hotel/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")

		if authToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Auth token is empty")
		}

		if err := helpers.ValidateToken(authToken); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth token")
		}

		return next(c)
	}
}
