package routers

import (
	"github.com/labstack/echo"
)

func loginCode(c echo.Context) error {
	// from Configuration of this project
	/*
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
		return err*/
	return nil
}

func addUserInfo(openID, avatarURL, gender, nickname string) error {
	//dao.FindByOpenid
	return nil
}
