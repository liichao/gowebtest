package router

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	// 	viewContext["reverse"] = c.Echo().Reverse
	// }
	if debug {
		t.templates = template.Must(template.ParseFiles("./views/login.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// 定义一个全局变量
var renderer = &TemplateRenderer{
	//templates: template.Must(template.ParseGlob("./views/*.html")),
	templates: template.Must(template.ParseFiles("./views/login.html")),
}
