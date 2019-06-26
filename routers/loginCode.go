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
	_gender := c.QueryParam("gender")
	var gender int

	if _gender == "1" {
		gender = 1
	} else {
		gender = 0
	}

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
	err = addUserInfo(openID, avatarURL, nickname, gender)
	if err != nil {
		fmt.Print("cannot add userInfo by openID: %s, error: %s/n", openID, err)
		return c.String(http.StatusOK, "")
	}

	var R struct{ OpenID, SessionKey string }
	R.OpenID = openID
	R.SessionKey = sessionKey

	return c.JSON(http.StatusOK, R)
}

func addUserInfo(openID, avatarURL, nickname string, gender int) error {
	d := dao.NewDaoUserInfo()
	isHave := d.HaveOpenid(openID)
	if isHave {
		fmt.Printf("Have this record of openid: %s", openID)
		return nil
	}

	userInfo := &dao.UserInfo{}
	userInfo.Openid = openID
	userInfo.AvatarUrl = avatarURL
	userInfo.Gender = gender
	userInfo.Nickname = nickname
	d.Insert(userInfo)

	return nil
}
