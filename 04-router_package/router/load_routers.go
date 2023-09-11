package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/04-router_package/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)
}
