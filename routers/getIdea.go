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

	df := dao.NewDaoFriend()
	defer df.Close()
	friends, err := df.FindByObjectAndStatus(Openid, 1)
	if err != nil {
		friends = nil
	}

	for _, item := range list {
		u := &jsonType{}
		u.Iid = item.Iid
		u.Openid = item.Openid
		u.Time = item.Time
		u.Content = item.Content
		u.Like = item.Like
		u.Favorite = item.Favorite

		user, friend := checkAnonymous(u.Openid, friends)
		u.AvatarUrl = user.AvatarUrl
		u.Gender = user.Gender
		u.Nickname = user.Nickname
		u.Friend = friend

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
		return c.JSON(http.StatusBadRequest, "error : Page should be number")
	}
	d := dao.NewDaoIdea()
	defer d.Close()
	list, err := d.QueryPage(page)
	json := convertType(list, Openid)
	return c.JSON(http.StatusOK, json)
}
