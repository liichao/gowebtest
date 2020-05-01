package main

import (
	_ "github.com/liichao/gowebtest/docs"
	"github.com/liichao/gowebtest/router"
)

// @title Swagger Example API
// @version 1.0
// @description 简易api描述文档.
// @host localhost:801
// @BasePath /
func main() {
	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//		return c.String(http.StatusOK, "Hello, World!")
	//	})
	//	e.Logger.Fatal(e.Start(":3123"))
	router.Run()
}
