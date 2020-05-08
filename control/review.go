package control

import (
	"../model"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ArticleReviewAll 查询所有评论
func ArticleReviewAll(ctx echo.Context) error {
	modes, err := model.ArticleReviewAll()
	if err != nil {
		return ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("所有数据", modes))
	//return ctx.JSONPretty(http.StatusOK, modes, " ")
}
