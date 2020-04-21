package main

import "./router"

func main() {
	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//		return c.String(http.StatusOK, "Hello, World!")
	//	})
	//	e.Logger.Fatal(e.Start(":3123"))
	router.Run()
}
