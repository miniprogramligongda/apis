package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysql_connect() *DB {

	// root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8
	// root:Mysql的密码@tcp(ip地址与端口)/数据库名称?charset=utf8
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		print("f")
		panic(err)
	}
	rows, _ := db.Query("desc user_info")
	defer rows.Close()
	dbNames := make([]string, 0)
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return
		}
		dbNames = append(dbNames, name)
	}
	fmt.Print(dbNames)

	stmt, _ := db.Prepare("update user_info set username=? where id=?")
	res, _ := stmt.Exec("lisi", 1)
	affect, _ := res.RowsAffected()
	fmt.Print(affect)
	//defer db.Close()
	return db
}
