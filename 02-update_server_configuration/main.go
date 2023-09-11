package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	//import server package
	"github.com/ujblockchain/echobootcamp/02-update_server_configuration/server"
)

func main() {
	//init echo
	app := echo.New()

	//add context
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//start server;
	server.SetServer(app)

}
