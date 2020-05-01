package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//User 表
type User struct {
	ID         int       `json:"id" db:"USER_ID"`
	User       string    `json:"user" db:"USER_NAME"`
	PassWord   string    `json:"password" db:"PASSWD"`
	Phone      string    `json:"phone" db:"MOBILE_NUM"`
	Email      string    `json:"email" db:"EMAIL"`
	Status     int       `json:"status" db:"STATE"`
	CreateTime time.Time `json:"create_time" db:"CREATE_TIME"`
}

// UserLogin 用户登陆
func UserLogin(user string) (User, error) {

	mod := User{}
	err := Db.Get(&mod, "select * from user WHERE USER_NAME=? LIMIT 1;", user)
	return mod, err
}

// UserToken Jwt info
type UserToken struct {
	ID       int
	PassWord string
	User     string
	jwt.StandardClaims
}
