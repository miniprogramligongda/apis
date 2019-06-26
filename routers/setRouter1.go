package routers

import (
	"github.com/labstack/echo"
)

// SetRouter1 :
// Be modified by NikoKVCS
func SetRouter1(e *echo.Echo) {

	// For example
	e.GET("getSearchUser", getSearchUser)
	e.POST("postUserinfo", postUserinfo)
	e.POST("postIdea", postIdea)
}
