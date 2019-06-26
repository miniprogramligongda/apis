package routers

import (
	"net/http"
	"strconv"

	"../dao"
	"github.com/labstack/echo"
)

func checkFriends(openid string, list []*dao.Idea) {

}

func getIdea(c echo.Context) error {
	Openid := c.QueryParam("Openid")
	PageRaw := c.QueryParam("Page")
	page, err := strconv.Atoi(PageRaw)
	if err != nil {
		return c.JSON(http.StatusOK, "error : Page should be number")
	}
	d := dao.NewDaoIdea()
	list, err := d.QueryPage(page)
	checkFriends(Openid, list)
	return c.JSON(http.StatusOK, list)
}
