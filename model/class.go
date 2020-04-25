package model

import (
	"errors"
	"log"
)

// Class 类别
type Class struct {
	ID   int64  `db:"ID,omitempty"`
	Name string `db:"Name,omitempty"`
	Desc string `db:"Desc,omitempty"`
}

// ClassPage 分页
func ClassPage(pi, ps int) ([]Class, error) {
	modes := make([]Class, 0, ps)
	err := Db.Select(&modes, "select * from class limit ?,?", (pi-1)*ps, ps)
	return modes, err
}

// ClassCount 查询所有总数
func ClassCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from class")
	return count
}

// ClassAll 查询所有
func ClassAll() ([]Class, error) {
	modes := make([]Class, 0, 8)
	err := Db.Select(&modes, "select * from class")
	log.Println("errr:", err)
	log.Println(modes)
	return modes, err
}

// ClassGet 查询某一条
func ClassGet(id int64) (*Class, error) {
	modes := &Class{}
	err := Db.Get(modes, "select * from class where id =? limit 1", id)
	return modes, err
}

// 添加

// ClassAdd 添加分类
func ClassAdd(mod *Class) error {
	tx, _ := Db.Begin()
	result, err := tx.Exec("insert into class(`Name`,`Desc`) values(?,?)", mod.Name, mod.Desc)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows 小于1")
	}
	tx.Commit()
	return nil
}

// 修改
// 删除