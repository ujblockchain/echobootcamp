package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/04-router_package/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
