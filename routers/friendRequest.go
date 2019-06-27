package routers

import (
	"../dao"
	"github.com/labstack/echo"
	"net/http"
)

type fRequest struct {
	ObjectOpenid  string
	SubjectOpenid string
	Notes         string
}

func friendRequest(c echo.Context) error {
	fr := &fRequest{}
	err := c.Bind(fr)

	friend := &dao.Friend{}
	friend.ObjectOpenid = fr.ObjectOpenid
	friend.SubjectOpenid = fr.SubjectOpenid
	friend.Notes = fr.Notes
	friend.Status = 0

	d := dao.NewDaoFriend()
	err = d.Insert(friend)
	if err != nil {
		return c.String(http.StatusAccepted, "request duplicate.") // TODO: 加一个 IsExist（O,S）
	}

	return c.String(http.StatusOK, "request succeed.")
}
