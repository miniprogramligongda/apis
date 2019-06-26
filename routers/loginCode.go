package routers

import (
	"../conf"
	"../dao"
	"fmt"
	"github.com/labstack/echo"
	"github.com/medivhzhan/weapp"
	"net/http"
)

// @appID 小程序 appID
// @secret 小程序的 app secret
// @code 小程序登录时获取的 code
// res, err := weapp.Login(appID, secret, code)
// if err != nil {
//     // handle error
// }

// // res.OpenID
// // res.SessionKey
// // res.UnionID
// fmt.Printf("返回结果: %#v", res)

func loginCode(c echo.Context) error {
	// from Configuration of this project
	appID := conf.Conf.AppID
	secret := conf.Conf.Secret

	code := c.QueryParam("code")
	avatarUrl := c.QueryParam("avatarUrl")
	nickname := c.QueryParam("nickname")
	gender := c.QueryParam("gender")

	res, err := weapp.Login(appID, secret, code)
	if err != nil {
		fmt.Errorf("%s", err)
		return c.String(http.StatusOK, "getOpenIDFailed")
	}

	openID := res.OpenID

}

func addUserInfo(openID, avatarUrl, gender, nickname string) error {
	dao.Insert
}
