package router

import (
	"../control"
	"github.com/labstack/echo"
)

var debug = true

// Run app start
func Run() {
	app := echo.New()
	app.HideBanner = true           // 隐藏banner
	app.Renderer = renderer         // 注册模板
	app.Static("/static", "static") // 加载静态文件目录
	app.Static("/views", "views")   // 加载静态文件目录
	app.GET("/", control.Index)
	app.GET("/login", control.LoginView)

	dzero := app.Group("dzero", ServerHeader) // 所有以dzero开头的页面都需要token
	AdmRouter(dzero)
	api := app.Group("/api") //  api 接口组
	APIRouter(api)
	app.Start(":13123")
}
