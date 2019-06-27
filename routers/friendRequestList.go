package routers

import (
	"../dao"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type fRequestList struct {
	SubjectOpenid string
	Notes         string
}

func friendRequestList(c echo.Context) error {
	objectOpenid := c.QueryParam("ObjectOpenid")
	d := dao.NewDaoFriend()
	defer d.Close()

	list, err := d.FindBySubjectAndStatus(objectOpenid, 0)
	if err != nil {
		fmt.Println(err)
	}

	result := make([]*fRequestList, 0)
	for _, item := range list {
		f := &fRequestList{}
		f.SubjectOpenid = item.SubjectOpenid
		f.Notes = item.Notes
		result = append(result, f)
	}

	return c.JSON(http.StatusOK, result)
}
