package routers

import (
	"../dao"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type frList struct {
	ObjectOpenid  string
	SubjectOpenid string
}

func friendRequestAgree(c echo.Context) error {
	frl := &frList{}
	err := c.Bind(frl)

	objectOpenid := frl.ObjectOpenid
	subjectOpenid := frl.SubjectOpenid

	d := dao.NewDaoFriend()
	err = d.UpdateStatusByOSbject(subjectOpenid, objectOpenid, 1)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusAccepted, "agree failed.")
	}

	newFriend := &dao.Friend{}
	newFriend.ObjectOpenid = objectOpenid
	newFriend.SubjectOpenid = subjectOpenid
	newFriend.Status = 1

	err = d.Insert(newFriend)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusAccepted, "agree failed.")
	}

	return c.String(http.StatusOK, "agree already.")
}
