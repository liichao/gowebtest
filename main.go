package main

import (
	_ "./docs"
	"./router"
)

// @title 在线api接口文档
// @version 1.0
// @name Dzer0
// @description 在线api接口文档
// @host 127.0.0.1:801
// @BasePath /
func main() {
	router.Run()
}
