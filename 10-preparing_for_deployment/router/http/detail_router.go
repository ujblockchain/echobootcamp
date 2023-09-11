package http

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/10-preparing_for_deployment/controller/context/pages"
)

func DetailsRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}
