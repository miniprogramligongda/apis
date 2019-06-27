package routers

import (
	"github.com/labstack/echo"
)

// SetRouter2 :
// Be modified by SmallXeon
func SetRouter2(e *echo.Echo) {

	// For example
	e.GET("getSearchUser", getSearchUser)

	//
	e.POST("loginCode", loginCode)

	// friend router
	e.POST("friendRequest", friendRequest)
	e.GET("friendRequestList", friendRequestList)
	e.POST("friendRequestAgree", friendRequestAgree)
	e.GET("friendList", friendList)

}
