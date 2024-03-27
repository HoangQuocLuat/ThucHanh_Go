package router

import (
	handler "thuchanh_go/handler/account"

	"github.com/gin-gonic/gin"
)

type API struct {
	Gin         *gin.Engine
	AccHandler  handler.AccountHandler
	// ChatHandler handler.UserChatHandler
}

func (api *API) SetupRoute() {
	api.Gin.POST("/user/register", api.AccHandler.RegisHandler())
	api.Gin.GET("/user/login", api.AccHandler.LoginHandler())
	// api.Gin.GET("/user/chat"), api.ChatHandler.
}
