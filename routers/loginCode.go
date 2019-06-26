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
	sessionKey := res.SessionKey

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("run time panic: %v", err)
		}
	}()
	err = addUserInfo(openID, avatarURL, gender, nickname)
	if err != nil {
		fmt.Print("cannot add userInfo by openID: %s, error: %s/n", openID, err)
		return c.String(http.StatusOK, "")
	}

	var R struct{ OpenID, SessionKey string }
	R.OpenID = openID
	R.SessionKey = sessionKey

	return c.JSON(http.StatusOK, R)
}

func addUserInfo(openID, avatarURL, gender, nickname string) error {
	d := dao.NewDaoUserInfo()
	isHave := d.HaveOpenid(openID)
	if isHave {
		err := fmt.Errorf("Have this record of openid: %s", openID)
		return nil
	}

	userInfo := &dao.UserInfo{}
	userInfo.OpenID = openID
	userInfo.avatarUrl = avatarURL
	userInfo.Gender = gender
	userInfo.Nickname = nickname
	d.Insert(userInfo)

	return nil
}
