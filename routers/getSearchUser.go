package routers

import (
	"github.com/labstack/echo"
	"net/http"
)

func getSearchUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
