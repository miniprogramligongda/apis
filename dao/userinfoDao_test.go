package dao

import (
	"fmt"
	"testing"
	"time"
)

// func TestUserinfoDao
//
func TestUserinfoDao(t *testing.T) {
	u := &UserInfo{}
	u.Openid = time.Now().Format("2006-01-02 15:04:05")
	u.AvatarUrl = time.Now().Format("2006-01-02")
	u.Gender = 1
	u.Nickname = "Niko"
	d := NewDaoUserInfo()
	//fmt.Printf("Prepare to insert : %v\n", u)
	d.HaveOpenid("shshs")
	result, _ := d.FindByOpenid(u.Openid)
	fmt.Printf("result : %v", result)
}
