package router

import (
	"../control"
	"github.com/labstack/echo"
)

var debug = true

// Run app start
func Run() {
	app := echo.New()
	app.Renderer = renderer         // 注册模板
	app.Static("/static", "static") // 加载静态文件目录
	app.Static("/views", "views")   // 加载静态文件目录
	app.GET("/", control.Index)
	app.GET("/login", control.LoginView)
	app.POST("/api/login", control.UserLogin)
	app.Start(":13123")
}
