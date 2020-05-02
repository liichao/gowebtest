package router

import (
	"../control"

	"github.com/labstack/echo"
)

// APIRouter API访问接口路由地址
func APIRouter(api *echo.Group) {
	api.POST("/login", control.UserLogin)
	api.GET("/class/all", control.ClassAll)
	api.GET("/class/page/:pi/:ps", control.ClassPage)
	api.GET("/class/get/:id", control.ClassGet)
	api.GET("/article/all", control.ArticleAll)               // 查询所有文章
	api.GET("/article/get/:id", control.ArticleGet)           // 查看一篇文章  http://url/api/article/get/文章id
	api.GET("/article/:cid/:pi/:ps", control.ArticleGetClass) // 查看一篇文章  http://url/api/article/文章分类id/页数/一页显示数量
}
