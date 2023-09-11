package main

import (
	"github.com/labstack/echo/v4"
	//import server package
	"github.com/ujblockchain/echobootcamp/03-using_context_package/controller/context/pages"
	"github.com/ujblockchain/echobootcamp/03-using_context_package/server"
)

func main() {
	//init echo
	app := echo.New()

	//add context
	app.GET("/", pages.IndexContext)

	//start server;
	server.SetServer(app)

}
