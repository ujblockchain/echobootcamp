package main

import (
	"github.com/labstack/echo/v4"
	//import server package
	"github.com/ujblockchain/echobootcamp/02-update_server_configuration/server"
	"github.com/ujblockchain/echobootcamp/03-using_context_package/controller/context/pages"
)

func main() {
	//init echo
	app := echo.New()

	//add context
	app.GET("/", pages.IndexContext)

	//start server;
	server.SetServer(app)

}
