package main

import (
	"github.com/labstack/echo/v4"
	//import server package
	"github.com/ujblockchain/echobootcamp/05-template_static_files/constant"
	"github.com/ujblockchain/echobootcamp/05-template_static_files/router"
	"github.com/ujblockchain/echobootcamp/05-template_static_files/server"
)

func main() {
	//init echo
	app := echo.New()

	//load static files
	constant.LoadStatic(app)

	//load template folder
	app.Renderer = constant.LoadTemplate()

	//add context
	router.LoadAllRouters(app)

	//start server;
	server.SetServer(app)

}
