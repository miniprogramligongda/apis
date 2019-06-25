package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"./dao"
	"./util"
)

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func TestMysql(t *testing.T) {
	db := dao.MysqlConstruct()
	defer db.Disconnect()

	id := RandInt(5, 1000)
	username := time.Now().Format("2006-01-02 15:04:05")
	password := time.Now().Format("2006-01-02")

	db.Insert("INSERT user_info SET id=?,username=?,password=?", id, username, password)
	fmt.Printf("Insert test : %d %s %s\n", id, username, password)

	rows := db.Query("SELECT * FROM user_info WHERE id=?", id)
	for rows.Next() {
		var _id int
		var _username string
		var _password string
		err := rows.Scan(&_id, &_username, &_password)
		util.CheckErr(err)
		fmt.Printf("Query test : %d %s %s\n", _id, _username, _password)
	}

	update_username := "nikokvcs"
	db.Update("update user_info set username=? where id=?", update_username, id)
	fmt.Printf("Update test : %d %s\n", id, update_username)

	rows = db.Query("SELECT * FROM user_info WHERE id=?", id)
	for rows.Next() {
		var _id int
		var _username string
		var _password string
		err := rows.Scan(&_id, &_username, &_password)
		util.CheckErr(err)
		fmt.Printf("Query test : %d %s %s\n", _id, _username, _password)
		if _username != update_username {
			t.Errorf("Username updated failed")
		}
	}

	fmt.Printf("Delete test : %d\n", id)
	db.Delete("delete from user_info where id=?", id)
	rows = db.Query("SELECT * FROM user_info WHERE id=?", id)
	if rows.Next() {
		t.Errorf("Deleted failed")
	}

}
