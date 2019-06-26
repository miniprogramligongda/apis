package routers

import (
	"net/http"

	"../dao"
	"../util"
	"github.com/labstack/echo"
)

func postUserinfo(c echo.Context) error {
	d := dao.NewDaoUserInfo()
	defer d.Close()
	/*
	   	var config ConfigStruct
	       if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
	           fmt.Println("================json str 转struct==")
	           fmt.Println(config)
	           fmt.Println(config.Host)
	       }
	*/
	u := &dao.UserInfo{}
	err := c.Bind(u)
	util.CheckErr(err)
	_, err = d.FindByOpenid(u.Openid)
	if err != nil {
		d.Insert(u)
	}

	d.Close()
	return c.String(http.StatusOK, "ok")
}
