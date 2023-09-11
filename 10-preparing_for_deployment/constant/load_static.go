package constant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	//get os path
	path, _ := os.Executable()

	// get file path without the app added to it
	filePath := filepath.Dir(path)

	//template location
	staticFolder := fmt.Sprintf("%v/repository/assets", filePath)


	//get path to static file like css, js fonts, images etc
	app.Static("static", staticFolder)
}
