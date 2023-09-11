package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/08-database_using_gorm/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)

	//form router
	http.FormRouter(app)
}
