package router

import (
	"../control"
	"github.com/labstack/echo"
)

// AdmRouter 后台登陆页面
func AdmRouter(api *echo.Group) {
	api.GET("/", control.AdminIndexView)
}
