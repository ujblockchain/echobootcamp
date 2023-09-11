package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/09-details_context_router_form_redirect/controller/context/pages"
)

func FormRouter(app *echo.Echo) {
	app.POST("/record", pages.FormContext)
}
