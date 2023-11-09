package authR

import (
	"thuchanh_go/handler/authH"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	userRoutes := r.Group("auth")
	{
		userRoutes.POST("/register", authH.RegisterHandler())
		userRoutes.POST("/login", authH.LoginHandler())
	}
}
