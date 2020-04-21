package control

import (
	"log"

	"../model"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

type login struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
}

func UserLogin(ctx echo.Context) error {
	postInfo := login{}
	err := ctx.Bind(&postInfo)
	log.Println(postInfo)
	if err != nil {
		//return ctx.String(200, err.Error())
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.UserLogin(postInfo.User)
	log.Println(postInfo.PassWord)
	log.Println(mod.PassWord)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名或密码不正确", err.Error()))
	}

	if mod.PassWord != postInfo.PassWord {
		return ctx.JSON(utils.ErrIpt("用户名或密码不正确"))
	}
	return ctx.JSON(utils.Succ("登陆成功", mod))
	//return ctx.JSONP(http.StatusOK, "", &postInfo)
}
