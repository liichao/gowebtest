package control

import (
	"github.com/labstack/echo"
)

// Index info
func Index(ctx echo.Context) error {

	return ctx.String(200, "hello,GoLang")
}
