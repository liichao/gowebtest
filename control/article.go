package control

import (
	"log"
	"strconv"

	"github.com/liichao/gowebtest/model"

	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ArticleAll 接口
// @Summary 查询所有文章
// @Description 获取article表中的所有数据，慎用
// @Produce  json
// @Success 200
// @Failure 500
// @Router /api/article/all [get]
func ArticleAll(ctx echo.Context) error {
	modes, err := model.ArticleAll()
	if err != nil {
		return ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("所有数据", modes))
	//return ctx.JSONPretty(http.StatusOK, modes, " ")
}

// ArticleGet 获取一篇文章信息
// @Summary 获取一篇文章信息
// @Description 打开一篇文章，并获取信息
// @Produce  json
// @Param id path int true "ARTICLE_ID" 文章id
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string	"error"
// @Router /api/article/get/{id} [get]
func ArticleGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	//id, err := strconv.ParseInt("2", 10, 64) // 获取id,按照10进制 转换成int64
	log.Println(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return ctx.JSON(utils.ErrIpt("ID输入错误"))
	}
	mod, err := model.ArticleGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到结果"))
	}
	return ctx.JSON(utils.Succ("data", mod))
}
