package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/09-details_context_router_form_redirect/controller/context/pages"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}
