package main

import (
	"thuchanh_go/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// init server
	r := gin.Default()

	// router
	r.POST("/user/login", handler.LoginUserHandler())
	r.POST("/user/register", handler.PostUserHandler())
	r.PUT("user/upacc/:user_id", handler.PutUserHandler())
	// run server
	r.Run("127.0.0.1:8888")
}
