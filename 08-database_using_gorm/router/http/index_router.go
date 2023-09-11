package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
