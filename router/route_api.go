package router

import (
	"../control"

	"github.com/labstack/echo"
)

// APIRouter API访问接口路由地址
func APIRouter(api *echo.Group) {
	api.POST("/login", control.UserLogin)
	api.GET("/class/all", control.ClassAll)
	api.GET("/class/page", control.ClassPage)
}
