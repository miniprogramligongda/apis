package routers

import (
	"net/http"
	"strconv"

	"../dao"

	"github.com/labstack/echo"
)

func getUnfav(c echo.Context) error {
	Openid := c.QueryParam("Openid")
	Iid := c.QueryParam("Iid")
	iid, err := strconv.ParseInt(Iid, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "error : tell h1astro Iid should be numeric")
	}
	d := dao.NewDaoIdea()
	defer d.Close()
	d.DecrementFavs(iid)

	dFav := dao.NewDaoFavorite()
	defer dFav.Close()
	dFav.DelFavOfUser(Openid, iid)
	return c.String(http.StatusOK, "hair++")
}
