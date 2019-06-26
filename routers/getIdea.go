package routers

import (
	"net/http"
	"strconv"
	"time"

	"../dao"
	"github.com/labstack/echo"
)

type jsonType struct {
	Iid       int64
	Openid    string
	Time      time.Time
	Content   string
	Like      int64
	Favorite  int64
	AvatarUrl string
	Gender    int
	Nickname  string
	Friend    int
}

func convertType(list []*dao.Idea, Openid string) []*jsonType {
	result := make([]*jsonType, 0)
	for _, item := range list {
		u := &jsonType{}
		u.Iid = item.Iid
		u.Openid = item.Openid
		u.Time = item.Time
		u.Content = item.Content
		u.Like = item.Like
		u.Favorite = item.Favorite

		d := dao.NewDaoUserInfo()
		defer d.Close()
		user, err := d.FindByOpenid(u.Openid)
		if err == nil && user != nil {
			u.AvatarUrl = user.AvatarUrl
			u.Gender = user.Gender
			u.Nickname = user.Nickname
		}

		checkFriends(u, Openid)
		if u.Friend == 0 {
			u.AvatarUrl = "https://gss0.baidu.com/-vo3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D600%2C800/sign=e746333c8cd4b31cf0699cbdb7e60b47/d788d43f8794a4c2d0b5cf5409f41bd5ad6e393e.jpg"
		}
		result = append(result, u)
	}
	return result
}

func checkFriends(u *jsonType, Openid string) {
	u.Friend = 0
}

func getIdea(c echo.Context) error {
	Openid := c.QueryParam("Openid")
	PageRaw := c.QueryParam("Page")
	page, err := strconv.Atoi(PageRaw)
	if err != nil {
		return c.JSON(http.StatusOK, "error : Page should be number")
	}
	d := dao.NewDaoIdea()
	defer d.Close()
	list, err := d.QueryPage(page)
	json := convertType(list, Openid)
	return c.JSON(http.StatusOK, json)
}
