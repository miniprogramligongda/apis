package main

import (
	"crypto/md5"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	db := MysqlConstruct()
	defer db.Disconnect()

	timeStr := time.Now().Format("2006-01-02 15:04:05")

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(timeStr))
	username := Md5Inst.Sum([]byte(""))

	username :=
		db.Insert("INSERT user_info SET id=?,username=?,password=?")
	t.Errorf("uppercase(%s) = %s,must be %s", ut.in, uc, ut.out)
}
