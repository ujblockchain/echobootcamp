package main

import (
	"github.com/labstack/echo/v4"
	//import constant package
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/constant"
	//import router package
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/router"
	//import server package
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/server"
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
