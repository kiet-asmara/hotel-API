package helpers

import (
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (*JWTClaim, error) {
	token := c.Request().Header.Get("Authorization")
	claims, err := DecodeToken(token)
	if err != nil {
		return &JWTClaim{}, err
	}
	return claims, nil
}
