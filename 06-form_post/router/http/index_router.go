package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/06-form_post/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
