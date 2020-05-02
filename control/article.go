package control

import (
	"log"
	"strconv"

	"../model"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ArticleAll 查询所有分类
func ArticleAll(ctx echo.Context) error {
	modes, err := model.ArticleAll()
	if err != nil {
		return ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("所有数据", modes))
	//return ctx.JSONPretty(http.StatusOK, modes, " ")
}

// ArticleGet 查询一条分类
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

// ArticleGetClass 分类分页
func ArticleGetClass(ctx echo.Context) error {
	// pi 获取分页
	pi, err := strconv.Atoi(ctx.Param("pi"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页索引错误"))
	}
	if pi < 1 {
		return ctx.JSON(utils.ErrIpt("分页索引范围错误"))
	}
	// ps 获取每页显示数量
	ps, err := strconv.Atoi(ctx.Param("ps"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页大小错误"))
	}
	// 设定一页内容不得少于1 或大于50 如果超出则修改成6
	if ps < 1 || ps > 50 {
		ps = 6
	}
	// 获取文章分类id,并判断文章分类是否存在
	classID, err := strconv.ParseInt(ctx.Param("cid"), 10, 64)
	if err != nil {
		log.Println(err)
		return ctx.JSON(utils.ErrIpt("文章分类ID错误"))
	}
	classInfo, err := model.ClassGet(classID)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("文章分类不存在"))
	}
	log.Println(classInfo.Name)
	// 获取文章总数
	count := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	pageInfo, err := model.ArticleGetClass(pi, ps, classInfo.ID)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	return ctx.JSON(utils.Page("data", pageInfo, count))
}
