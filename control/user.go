package control

import (
	"time"

	"github.com/liichao/gowebtest/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

type login struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
}

// UserLogin 登陆判断
// @Summary 用户登陆
// @Description 判断用户是否登陆成功
// @Produce  json
// @Accept  json
// @Success 200
// @Failure 500
// @Router /login [post]
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
			// ExpiresAt: 15000,// 过期时间
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // 取登陆时间+2小时后过期
		},
	}
	usertoken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	ss, err := usertoken.SignedString([]byte("2599"))
	// fmt.Printf("%v,%v,%v", mod.ID, ss, err)
	return ctx.JSON(utils.Succ("登陆成功", ss))
	//return ctx.JSONP(http.StatusOK, "", &postInfo)
}
