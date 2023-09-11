package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexContext(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
