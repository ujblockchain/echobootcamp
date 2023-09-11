package main

import (
	"github.com/labstack/echo/v4"
	//import server package
	"github.com/ujblockchain/echobootcamp/04-router_package/router"
	"github.com/ujblockchain/echobootcamp/04-router_package/server"
)

func main() {
	//init echo
	app := echo.New()

	//add context
	router.LoadAllRouters(app)

	//start server;
	server.SetServer(app)

}
