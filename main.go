package main

import (
	"time"

	"./conf"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Server.WriteTimeout = 20 * time.Minute
	e.Server.ReadTimeout = 20 * time.Minute

	//e.GET("searchuser", getSearchUser)
	setRouter(e)
	e.Logger.Fatal(e.StartTLS(":1323", conf.Conf.CertFile, conf.Conf.KeyFile))
}
