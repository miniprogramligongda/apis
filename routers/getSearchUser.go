package getSearchUser

import (
	"net/http"

	"github.com/labstack/echo"
)

func getSearchUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
