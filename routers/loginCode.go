package routers

import (
	"../conf"
	"../dao"
	"fmt"
	"github.com/labstack/echo"
	"github.com/medivhzhan/weapp"
	"net/http"
)

func loginCode(c echo.Context) error {
	// from Configuration of this project
	appID := conf.Conf.AppID
	secret := conf.Conf.Secret

	code := c.QueryParam("code")
	avatarURL := c.QueryParam("avatarUrl")
	nickname := c.QueryParam("nickname")
	gender := c.QueryParam("gender")

	res, err := weapp.Login(appID, secret, code)
	if err != nil {
		fmt.Errorf("%s", err)
		return c.String(http.StatusOK, "getOpenIDFailed")
	}

	openID := res.OpenID
	addUserInfo(appID, avatarURL, gender, nickname)
	return err
}

func addUserInfo(openID, avatarURL, gender, nickname string) error {
	dao.FindByOpenid
	return nil
}
