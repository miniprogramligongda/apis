package routers

import (
	"net/http"

	"../dao"

	"github.com/labstack/echo"
)

func getFavList(c echo.Context) error {
	Openid := c.QueryParam("Openid")

	dFav := dao.NewDaoFavorite()
	defer dFav.Close()
	iidlist := dFav.GetFavListOfUser(Openid)

	jsons := make([]*dao.Idea, 0)
	di := dao.NewDaoIdea()
	defer di.Close()
	for _, v := range iidlist {
		idea := di.FindByIid(v)
		if idea != nil {
			jsons = append(jsons, idea)
		}
	}

	return c.JSON(http.StatusOK, jsons)
}
