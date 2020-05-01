package model

import "log"

// Article 类别
type Article struct {
	ID           int64  `db:"ARTICLE_ID,omitempty"`
	Class        string `db:"CLASS_ID,omitempty"`
	Title        string `db:"TITLE,omitempty"`
	Content      string `db:"CONTENT,omitempty"`
	BaiDuDown    string `db:"BAIDU_DOWN,omitempty"`
	BaiDuDownPwd string `db:"BAIDU_DOWN_PWD,omitempty"`
	TianYiDown   string `db:"TIANYI_DOWN,omitempty"`
	TianYiPwd    string `db:"TIANYI_DOWN_PWD,omitempty"`
	Status       string `db:"STATE,omitempty"`
}

// ArticleAll 查询所有
func ArticleAll() ([]Article, error) {
	modes := make([]Article, 0, 8)
	err := Db.Select(&modes, "select * from article")
	log.Println("errr:", err)
	log.Println(modes)
	return modes, err
}

// ArticleGet 查询某一条
func ArticleGet(id int64) (*Article, error) {
	modes := &Article{}
	err := Db.Get(modes, "select * from article where ARTICLE_ID =? limit 1", id)
	log.Println(err)
	return modes, err
}
