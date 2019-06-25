package dao

import (
	"database/sql"

	"../conf"
	"../util"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

// MysqlConstruct
//
func MysqlConstruct() Database {

	// root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8
	// root:Mysql的密码@tcp(ip地址与端口)/数据库名称?charset=utf8
	db, err := sql.Open("mysql", conf.Conf.DbAddr)
	util.CheckErr(err)

	database := Database{db}
	return database
}

// 参数的定义：
// statement:CREATE TABLE example ( id integer, data varchar(32) )
// 例如 db.CreateTable("CREATE TABLE example ( id integer, data varchar(32) )")
func (this *Database) CreateTable(statement string, args ...interface{}) {
	_, err := this.db.Exec(statement, args...)
	util.CheckErr(err)
}

// 参数的定义：
// statement:SELECT * FROM userinfo
func (this *Database) Query(statement string, args ...interface{}) *sql.Rows {
	rows, err := this.db.Query(statement, args...)
	util.CheckErr(err)
	return rows
}

// 参数的定义：
// statement:"delete from userinfo where uid=?"
// args: id
// 例如 db.Delete("delete from userinfo where uid=?", id)
func (this *Database) Delete(statement string, args ...interface{}) {
	stmt, err := this.db.Prepare(statement)
	util.CheckErr(err)
	_, err = stmt.Exec(args...)
	util.CheckErr(err)
}

// 参数的定义：
// statement:update userinfo set username=? where uid=?
// args: "astaxieupdate", id
// 例如 db.Update("update userinfo set username=? where uid=?","astaxieupdate", id)
func (this *Database) Update(statement string, args ...interface{}) {
	stmt, err := this.db.Prepare(statement)
	util.CheckErr(err)
	_, err = stmt.Exec(args...)
	util.CheckErr(err)
}

// 参数的定义：
// statement:INSERT userinfo SET username=?,departname=?,created=?
// values: "astaxie", "研发部门", "2012-12-09"
// 例如 db.Insert("INSERT userinfo SET username=?,departname=?,created=?", "astaxie", "研发部门", "2012-12-09")
func (this *Database) Insert(statement string, values ...interface{}) {
	// 插入数据
	stmt, err := this.db.Prepare(statement)
	util.CheckErr(err)
	_, err = stmt.Exec(values...)
	util.CheckErr(err)
}

func (this *Database) Disconnect() {
	this.db.Close()
}
