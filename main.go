package main

import (
	"thuchanh_go/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// init server
	r := gin.Default()

	// router
	r.GET("/user/:user_id", handler.GetUserHandler())
	r.POST("/user/regis", handler.PostUserHandler())
	// run server
	r.Run("127.0.0.1:8888")
}
