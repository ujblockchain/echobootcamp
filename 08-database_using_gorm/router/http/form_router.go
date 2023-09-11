package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/controller/context/pages"
)

func FormRouter(app *echo.Echo) {
	app.POST("/record", pages.FormContext)
}
