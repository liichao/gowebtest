package router

import (
	"../control"
	"github.com/labstack/echo"
)

func Run() {
	app := echo.New()
	// 加载静态文件目录
	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", control.Index)
	app.POST("/api/login", control.UserLogin)
	app.Start(":13123")
}
