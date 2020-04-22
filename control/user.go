package control

import (
	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

type login struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
}

// UserLogin 登陆判断
func UserLogin(ctx echo.Context) error {
	postInfo := login{}
	err := ctx.Bind(&postInfo)
	// log.Println(postInfo)
	if err != nil {
		//return ctx.String(200, err.Error())
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.UserLogin(postInfo.User)
	// log.Println(postInfo.PassWord)
	// log.Println(mod.PassWord)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名或密码不正确", err.Error()))
	}

	if mod.PassWord != postInfo.PassWord {
		return ctx.JSON(utils.ErrIpt("用户名或密码不正确"))
	}
	UserClaims := model.UserToken{
		ID:       mod.ID,
		User:     mod.User,
		PassWord: mod.PassWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	usertoken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	ss, err := usertoken.SignedString([]byte("2599"))
	// fmt.Printf("%v,%v,%v", mod.ID, ss, err)
	return ctx.JSON(utils.Succ("登陆成功", mod, ss))
	//return ctx.JSONP(http.StatusOK, "", &postInfo)
}
