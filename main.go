package main

import (
	"./conf"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//e.GET("searchuser", getSearchUser)
	setRouter(e)
	e.Logger.Fatal(e.StartTLS(":1323", conf.Conf.CertFile, conf.Conf.KeyFile))
}
