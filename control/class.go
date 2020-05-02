package control

import (
	//"log"
	"log"
	"strconv"

	"../model"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ClassGet 查询一条分类
// @Summary 查询单个分类
// @Description 查询一个分类的详细信息
// @Produce  json
// @Param id path int true "分类ID"
// @Success 200  {string} string	"ok"
// @Failure 500
// @Router /api/class/get/{id} [get]
func ClassGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	//id, err := strconv.ParseInt("2", 10, 64) // 获取id,按照10进制 转换成int64
	log.Println(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return ctx.JSON(utils.ErrIpt("ID输入错误"))
	}
	mod, err := model.ClassGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到结果"))
	}
	return ctx.JSON(utils.Succ("data", mod))
}

// ClassAll 查询所有分类
// @Summary 获取所有分类
// @Description 获取所有分类
// @Produce  json
// @Success 200  {string} string	"ok"
// @Failure 500
// @Router /api/class/all [get]
func ClassAll(ctx echo.Context) error {
	modes, err := model.ClassAll()
	if err != nil {
		return ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", modes))
}

// ClassPage 分类分页
// @Summary  分类按分页获取
// @Description 分类按分页获取
// @Produce  json
// @Param pi path int true "当前页数"
// @Param ps path int true "一页显示多少数据"
// @Success 200  {string} string	"ok"
// @Failure 500
// @Router /api/class/page/{pi}/{ps} [get]
func ClassPage(ctx echo.Context) error {
	pi, err := strconv.Atoi(ctx.Param("pi"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页索引错误"))
	}
	if pi < 1 {
		return ctx.JSON(utils.ErrIpt("分页索引范围错误"))
	}
	ps, err := strconv.Atoi(ctx.Param("ps"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页大小错误"))
	}
	// 设定一页内容不得少于1 或大于50 如果超出则修改成6
	if ps < 1 || ps > 50 {
		ps = 6
	}
	count := model.ClassCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	pageInfo, err := model.ClassPage(pi, ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	return ctx.JSON(utils.Page("分类分页", pageInfo, count))
}

// ClassAdd 添加分类
func ClassAdd(ctx echo.Context) error {
	// {
	// 	"Name": "aaaa",
	// 	"Desc": "bbbb"
	// }
	ipt := &model.Class{}
	err := ctx.Bind(ipt)
	log.Println(ipt.Desc)
	log.Println(ipt.Name)
	if err != nil {
		log.Println(err)
		return ctx.JSON(utils.ErrIpt("输入内容有误"))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	err = model.ClassAdd(ipt)
	if err != nil {
		log.Println(err)
		return ctx.JSON(utils.ErrIpt("添加失败 "))
	}
	return ctx.JSON(utils.Succ("数据插入成功", ipt))
}
