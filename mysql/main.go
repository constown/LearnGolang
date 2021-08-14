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

func singleQuery(id int) {
	var u1 user
	// 1、查询单条记录的语句
	sqlStr := `select id,name,age from user where id=?;`
	// 从连接池拿一个连接出来去数据库查询单条数据
	// 拿到结果，一定要调用Scan方法,只有这个scan才会释放会连接池
	err := db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", u1)
}

func multipleQuery(n int) {
	//	1、sql语句
	sqlStr := `select id, name, age from user where id > ?;`
	// 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exex %s query failed, err: %v\n", sqlStr, err)
	}
	// 【重要】一定要关闭rows
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("关闭数据链接失败")
		}
	}(rows)
	//	循环取值
	for rows.Next() {
		var u1 user
		// 调用的是不同是Scan 所以一定要 close
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			return
		}
		fmt.Printf("%#v\n", u1)
	}
}

// 插入数据
func insert(name string, age int) {
	//	写sql语句
	sqlStr := `insert into user(name, age) values(?, ?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新数据
func update(newAge, id int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("updata failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("update rows:", n)
}

// 删除数据
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Println("删除了", n, "条数据")
}

// 预处理方式插入多条语句
func prepareInsert() {
	sqlStr := `insert into user(name, age) values (?, ?)`
	stmt, err := db.Prepare(sqlStr) // 先把sql语句发给mysql预处理
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	var m = map[string]int{
		"旺达":  30,
		"王相聚": 34,
		"涅卡":  54,
		"黑吧":  23,
		"林州":  12,
	}
	for k, v := range m {
		_, err := stmt.Exec(k, v) // 后续就只需要传值就可以了
		if err != nil {
			return
		}
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB err: %v\n", err)
	}
	fmt.Println("数据库连接成功！")
	insert("张大情", 89)
	//deleteRow(9)
	update(13, 5)
	singleQuery(2)
	multipleQuery(1)
	deleteRow(7)
	prepareInsert()
}
