package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//e.GET("searchuser", getSearchUser)
	setRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
