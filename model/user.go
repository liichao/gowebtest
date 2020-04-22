package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//User 表
type User struct {
	ID         int       `json:"id" db:"ID"`
	User       string    `json:"user" db:"User"`
	PassWord   string    `json:"password" db:"PassWord"`
	Phone      string    `json:"phone" db:"Phone"`
	Email      string    `json:"email" db:"Email"`
	Status     int       `json:"status" db:"Status"`
	CreateTime time.Time `json:"create_time" db:"CreateTime"`
}

// UserLogin 用户登陆
func UserLogin(user string) (User, error) {

	mod := User{}
	err := Db.Get(&mod, "select * from user WHERE user=? LIMIT 1;", user)
	return mod, err
}

// UserToken Jwt info
type UserToken struct {
	ID       int
	PassWord string
	User     string
	jwt.StandardClaims
}
