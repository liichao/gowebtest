package control

import (
	"net/http"

	"github.com/labstack/echo"
)

// LoginView 用户登陆页面
func LoginView(c echo.Context) error {
	// log.Println(c.Render(http.StatusOK, "hello", "World"))
	return c.Render(http.StatusOK, "login.html", nil)
}
