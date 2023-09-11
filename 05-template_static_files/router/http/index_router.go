package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/05-template_static_files/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
