package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	//init echo
	app := echo.New()

	//add context
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//start server; we ignore the error for now
	app.Start(":3000")

}
