package router

import (
	"html/template"
	"io"
	"log"

	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
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
		t.templates = template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// 定义一个全局变量
var renderer = &TemplateRenderer{
	//templates: template.Must(template.ParseGlob("./views/*.html")),
	templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html")),
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "盏茶阅读")
		tokenString := c.FormValue("token")
		log.Println(tokenString)
		clims := model.UserToken{}
		token, err := jwt.ParseWithClaims(tokenString, &clims, func(token *jwt.Token) (interface{}, error) {
			return []byte("2599"), nil
		})

		if err == nil && token.Valid {
			return next(c)
		} else {
			return c.JSON(utils.ErrJwt("Token过期"))
		}
	}
}
