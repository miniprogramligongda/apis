package main

import (
	"./routers"
	"github.com/labstack/echo"
	//"net/http"
)

func setRouter(e *echo.Echo) {
	routers.SetRouter1(e)
	routers.SetRouter2(e)
}
