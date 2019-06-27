package routers

import (
	"github.com/labstack/echo"
)

// SetRouter1 :
// Be modified by NikoKVCS
func SetRouter1(e *echo.Echo) {

	// For example
	e.GET("getSearchUser", getSearchUser)
	e.GET("getIdea", getIdea)
	e.POST("postUserinfo", postUserinfo)
	e.POST("postIdea", postIdea)
	e.GET("like", getLike)
	e.GET("fav", getFav)
	e.GET("unfav", getUnfav)
	e.GET("getFav", getFavList)
	e.POST("postComment", postComment)
	e.GET("getComment", getComment)
}
