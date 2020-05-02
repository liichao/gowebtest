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

// ArticleGet 打开一篇文章
func ArticleGet(id int64) (*Article, error) {
	modes := &Article{}
	err := Db.Get(modes, "select * from article where ARTICLE_ID =? limit 1", id)
	log.Println(err)
	return modes, err
}

// ArticleCount 查询所有总数
func ArticleCount() int {
	count := 0
	Db.Get(&count, "select count(ARTICLE_ID) as count from article")
	return count
}

// ArticleGetClass 分页
func ArticleGetClass(pi, ps int, classID int64) ([]Article, error) {
	// pi 第几页
	// ps 每页几条信息
	// classID 分类id
	modes := make([]Article, 0, ps)
	err := Db.Select(&modes, "select * from article where CLASS_ID=? limit ?,?", classID, (pi-1)*ps, ps)
	return modes, err
}
