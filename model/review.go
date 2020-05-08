package model

import (
	"log"
	"time"
)

// ArticleReview 文章评论
type ArticleReview struct {
	ARID          int64     `db:"AR_ID,omitempty"`
	ArticleID     int64     `db:"ARTICLE_ID,omitempty"`
	UserID        int64     `db:"USER_ID,omitempty"`
	OPTime        time.Time `db:"OP_DATE_TIME,omitempty"`
	ReviewComment string    `db:"REVIEW_COMMENT,omitempty"`
	Remark        string    `db:"REMARK,omitempty"`
}

// ArticleReviewAll 查询所有
func ArticleReviewAll() ([]ArticleReview, error) {
	modes := make([]ArticleReview, 0, 8)
	err := Db.Select(&modes, "select * from article_review")
	log.Println("errr:", err)
	log.Println(modes)
	return modes, err
}
