package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/05-template_static_files/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)
}
