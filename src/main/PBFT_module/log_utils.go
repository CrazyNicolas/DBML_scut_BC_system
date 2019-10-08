package PBFT_module

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db_username   = "root"
	db_password   = "root"
	db_ip_address = "127.0.0.1:3306"
	db_name       = "dbml"
	db_address    = db_username + ":" + db_password + "@tcp(" + db_ip_address + ")/" + db_name + "?charset=utf8"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/**
江声:
log的数据结构：
消息类型	v	n	d	i	m
*/

//这个函数就是用来写log文件用的 content数组其实里面装的应该是消息结构体利用 strcov 包里的函数转化成string
func log(content []string) {
	println(db_address)
	db, err := sql.Open("mysql", db_address)
	checkError(err)
	//stmt, err := db.Prepare("select * from test where id=?")
	//rows, err := stmt.Query("0")
	rows, err := db.Query("select * from test")
	checkError(err)
	for rows.Next() {
		var username string
		var passwd string
		var id string
		rows.Scan(&username, &passwd, &id)
		fmt.Println(username, passwd, id)
	}
}
