package utils

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// untuk log panic = internal server error
func LogError(c echo.Context, err error, stack []byte) error {
	log.Println(err)
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
