package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/06-form_post/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)

	//form router
	http.FormRouter(app)
}
