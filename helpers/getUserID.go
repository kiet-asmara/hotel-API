package helpers

import (
	"github.com/labstack/echo/v4"
)

func GetUserId(c echo.Context) (int, error) {
	token := c.Request().Header.Get("Authorization")
	claims, err := DecodeToken(token)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
