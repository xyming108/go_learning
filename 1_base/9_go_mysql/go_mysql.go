package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
Author: xym
Date: 2021/5/9 16:21
Project: myproject2_go_Learning
Description:
*/

//增
func insert(db *sql.DB) {
	insert, err := db.Exec("insert into user_info(username, password)values (?, ?)", "整了好久", "555")
	if err != nil {
		fmt.Println("插入失败！", err)
		return
	}
	affected, err := insert.RowsAffected()
	if err != nil {
		fmt.Println("失败！", err)
		return
	}
	fmt.Println("插入成功！", affected)
}

//删
func delete(db *sql.DB) {
	delete, err := db.Exec("delete from user_info where id = ?", 26)
	if err != nil {
		fmt.Println("删除失败！", err)
		return
	}
	affected, err := delete.RowsAffected()
	if err != nil {
		fmt.Println("失败！", err)
		return
	}
	if affected == 0 {
		fmt.Println("删除失败！", affected)
	} else if affected == 1 {
		fmt.Println("删除成功！", affected)
	}
}

//改
func update(db *sql.DB) {
	update, err := db.Exec("update user_info set username = ? where id = ?", "突然发现简单了", "31")
	if err != nil {
		fmt.Println("修改失败！", err)
		return
	}
	affected, err := update.RowsAffected()
	if err != nil {
		fmt.Println("失败！", err)
		return
	}
	fmt.Println("修改成功！", affected)
}

//查
func query(db *sql.DB) {
	query, err := db.Query("select * from user_info")
	if err != nil {
		fmt.Println("查找失败！", err)
		return
	}
	for query.Next() {
		var id int
		var username string
		var password string
		err := query.Scan(&id, &username, &password)
		if err != nil {
			fmt.Println("查找失败！", err)
			return
		}
		fmt.Println(id, username, password)
	}

	fmt.Println("查找成功！")
}

func main() {
	db, err := sql.Open("mysql", "root:564445@tcp(127.0.0.1:3306)/go_test")
	if err != nil {
		fmt.Println("打开失败！")
		return
	}
	defer db.Close()

	//insert(db)
	//delete(db)
	//update(db)
	query(db)
}
