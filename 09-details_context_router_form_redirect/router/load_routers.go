package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ujblockchain/echobootcamp/09-details_context_router_form_redirect/router/http"
)

func LoadAllRouters(app *echo.Echo) {
	//index router
	http.IndexRouter(app)

	//form router
	http.FormRouter(app)

	//details details
	http.DetailsRouter(app)
}
