package routers

import (
	"../dao"
	"../util"
	"github.com/labstack/echo"
)

type UserInfo struct {
	wxid      string
	avatarUrl string
	city      string
	country   string
	gender    int
	language  string
	nickName  string
	province  string
}

func postUserinfo(c echo.Context) error {
	u := new(UserInfo)
	err := c.Bind(u)
	util.CheckErr(err)
	db := dao.MysqlConstruct()
	db.Insert("INSERT user_info SET wxid=?,avatarUrl=?,city=?", u.wxid, u.avatarUrl, u.city)
	return nil
}
