package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/10-preparing_for_deployment/controller/context/pages"
)

func FormRouter(app *echo.Echo) {
	app.POST("/record", pages.FormContext)
}
