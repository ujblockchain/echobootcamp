package constant

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// init struct for template
type Template struct {
	templates *template.Template
}

// set render used for call template name eg index.html
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func LoadTemplate() *Template {
	//get os path
	path, _ := os.Executable()

	// get file path without the app added to it
	filePath := filepath.Dir(path)

	//template location
	templateFolder := fmt.Sprintf("%v/repository/templates/*", filePath)

	//get path to template folder
	template := &Template{
		templates: template.Must(template.ParseGlob(templateFolder)),
	}

	//return template path
	return template
}
