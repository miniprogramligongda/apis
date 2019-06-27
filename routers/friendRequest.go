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

	if fr.ObjectOpenid == fr.SubjectOpenid {
		return c.String(http.StatusBadRequest, "request self.")
	}

	friend := &dao.Friend{}
	friend.ObjectOpenid = fr.ObjectOpenid
	friend.SubjectOpenid = fr.SubjectOpenid
	friend.Notes = fr.Notes
	friend.Status = 0

	d := dao.NewDaoFriend()

	if have, _ := d.HaveFriendByOSbject(friend.ObjectOpenid, friend.SubjectOpenid); have {
		return c.String(http.StatusConflict, "request duplicate.")
	}

	err = d.Insert(friend)
	if err != nil {
		return c.String(http.StatusInternalServerError, "database status wrong.")
	}

	return c.String(http.StatusOK, "request succeed.")
}
