package routers

import (
	"../dao"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type fList struct {
	SubjectOpenid string
}

func friendList(c echo.Context) error {
	objectOpenid := c.QueryParam("ObjectOpenid")
	d := dao.NewDaoFriend()
	defer d.Close()

	list, err := d.FindByObjectAndStatus(objectOpenid, 1)
	if err != nil {
		fmt.Println(err)
	}

	result := make([]*fList, 0)
	for 

}
