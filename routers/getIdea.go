package routers

import (
	"net/http"
	"strconv"

	"../dao"
	"github.com/labstack/echo"
)

func getIdea(c echo.Context) error {
	Openid := c.QueryParam("Openid")
	PageRaw := c.QueryParam("Page")
	page, err := strconv.Atoi(PageRaw)
	if err != nil {
		return nil
	}
	d := dao.NewDaoIdea()
	list, err := d.QueryPage(page)
	if err == nil && len(list) != 0 {

	}

	return c.String(http.StatusOK, "Hello, World!")
}
