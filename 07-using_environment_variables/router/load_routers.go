package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/07-using_environment_variables/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)

	//form router
	http.FormRouter(app)
}
