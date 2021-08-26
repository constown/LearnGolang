package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// go 连接数据库

var db *sql.DB // 连接词池对象

func initDB() (err error) {
	//	连接数据库
	dsn := "root:root123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // 格式不正确
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库链接池最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲链接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

func transaction() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin transaction failed, err:%v\n", err)
		return
	}
	// 执行多个sql操作
	sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update user set age=age+2 where id=2`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		return
	}
	//	都执行成功之后，提交事务
	err = tx.Commit()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB err: %v\n", err)
	}
	fmt.Println("数据库连接成功！")
	transaction()
}
