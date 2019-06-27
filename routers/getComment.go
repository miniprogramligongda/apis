package routers

import (
	"net/http"
	"strconv"
	"time"

	"../dao"

	"github.com/labstack/echo"
)

type jsonCommentType struct {
	Iid       int64
	Openid    string
	Content   string
	Time      time.Time
	AvatarUrl string
	Gender    int
	Nickname  string
	Friend    int
}

func convertCommentType(raw []*dao.Comment, Openid string) []*jsonCommentType {
	result := make([]*jsonCommentType, 0)
	for _, item := range raw {
		u := &jsonCommentType{}
		u.Iid = item.Iid
		u.Openid = item.Openid
		u.Content = item.Content
		u.Time = item.Time

		user, friend := checkAnonymous(u.Openid, Openid)
		u.AvatarUrl = user.AvatarUrl
		u.Gender = user.Gender
		u.Nickname = user.Nickname
		u.Friend = friend

		result = append(result, u)
	}
	return result
}

func getComment(c echo.Context) error {
	Openid := c.QueryParam("Openid")
	Iid := c.QueryParam("Iid")
	iid, err := strconv.ParseInt(Iid, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "error : tell h1astro Iid should be numeric")
	}

	d := dao.NewDaoComment()
	defer d.Close()
	commentList := d.GetCommentsOfIdea(iid)
	jsons := convertCommentType(commentList, Openid)

	return c.JSON(http.StatusOK, jsons)
}
