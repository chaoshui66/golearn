package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUser(db *sql.DB) (user, error) {
	id := 1
	user := user{}
	err := db.QueryRow("SELECT id, name FROM `test` WHERE `id` = ?", id).Scan(&user.Id, &user.Name)
	// 此处只查询一行数据 返回一个struct 如果查询不到 应该返回一个error
	// 可以是原有error 的wrap 或者是自定义的error类型
	// 此处应当Wrap, 携带具体的user id, 或者原始的查询信息等
	if err != nil {
		newErr := fmt.Errorf("查询user时出错, ID 为 %d\n原始错误:\n%w", id, err)
		return user, newErr
	}
	return user, nil
}

func main() {
	db, err := sql.Open("mysql", "root:1221@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Println("数据库连接失败...")
	} else {
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				fmt.Println("关闭连接失败")
			}
		}(db)
		_, err = getUser(db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
