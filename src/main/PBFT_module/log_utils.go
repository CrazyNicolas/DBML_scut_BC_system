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
	db_name       = "dbml" //数据库的名称，在江声这里叫做dbml
	db_address    = db_username + ":" + db_password + "@tcp(" + db_ip_address + ")/" + db_name + "?charset=utf8"
)

/**
江声:
log的数据结构：
消息类型(enum)	v(int)	n(int)	d(varchar)	i(int)	m(json)
log为Replica的成员方法,传参数为各种message
*/

//这个函数就是用来写log文件用的 content数组其实里面装的应该是消息结构体利用 strconv 包里的函数转化成string
func (r *Replica) log(msg interface{}) {
	db, err := sql.Open("mysql", db_address)
	if err != nil {
		fmt.Println("打开数据库失败", err)
	}
	switch msg.(type) {
	case Prepreprare_Msg:
		//i在这里是-1，因为没有意义
		preprepare := msg.(Prepreprare_Msg)
		content := "insert into replica" + ToString(r.serialNumber) + " values(" + "'PRE-PREPARE'," + ToString(preprepare.v) +
			"," + ToString(preprepare.n) + "," + ToString(preprepare.digest) + ",-1," + ToString(preprepare.request) + ");"
		_, err := db.Exec(content)
		if err != nil {
			fmt.Println("添加preprepare到日志出错", err)
		}
	case Prepare_Msg:
		//m在这里是Null
		prepare := msg.(Prepare_Msg)
		content := "insert into replica" + ToString(r.serialNumber) + " values(" + "'PREPARE'," + ToString(prepare.v) +
			"," + ToString(prepare.n) + "," + ToString(prepare.digest) + "," + ToString(r.serialNumber) + ",null);"
		_, err := db.Exec(content)
		if err != nil {
			fmt.Println("添加prepare到日志出错", err)
		}
	case Commit_Msg:
		//m在这里也是Null
		commit := msg.(Commit_Msg)
		content := "insert into replica" + ToString(r.serialNumber) + " values(" + "'COMMIT'," + ToString(commit.v) +
			"," + ToString(commit.n) + "," + ToString(commit.digest) + "," + ToString(r.serialNumber) + ",null);"
		_, err := db.Exec(content)
		if err != nil {
			fmt.Println("添加commit到日志出错", err)
		}
	default:
	}
}
