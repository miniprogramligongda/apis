package routers

import (
	"time"

	"../dao"
	"../util"
	"github.com/labstack/echo"
)

type postidea struct {
	Openid  string
	Content string
}

func postIdea(c echo.Context) error {

	data := &postidea{}
	err := c.Bind(data)
	util.CheckErr(err)

	u := &dao.Idea{}
	//u.Iid = data.Openid + time.Now().Format("2006-01-02 15:04:05")
	u.Openid = data.Openid
	u.Time = time.Now()
	u.Content = data.Content
	u.Like = 0
	u.Favorite = 0

	d := dao.NewDaoIdea()
	defer d.Close()
	d.Insert(u)

	return nil
}
