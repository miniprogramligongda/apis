package routers

import (
	"net/http"
	"strconv"

	"../dao"
	"github.com/labstack/echo"
)

type postCommentData struct {
	Openid  string
	Iid     string
	Content string
}

func postComment(c echo.Context) error {
	u := &postCommentData{}
	err := c.Bind(u)
	if err != nil {
		c.String(http.StatusBadRequest, "hair-- please tell h1astro his post a wrong data")
	}
	iid, err := strconv.ParseInt(u.Iid, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "error : tell h1astro Iid should be numeric")
	}

	d := dao.NewDaoComment()
	defer d.Close()
	d.AddCommentToIdea(iid, u.Openid, u.Content)

	return c.String(http.StatusOK, "hair++")
}
