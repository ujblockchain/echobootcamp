package constant

import (
	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	//get path to static file like css, js fonts, images etc
	app.Static("static", "repository/assets")
}
