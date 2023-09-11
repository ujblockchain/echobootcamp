package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/07-using_environment_variables/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
