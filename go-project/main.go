package main

import (
	"thuchanh_go/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// init server
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// router
	r.POST("/user/login", handler.LoginUserHandler())
	r.POST("/user/register", handler.PostUserHandler())
	r.GET("user/infor/:user_id", handler.GetUserHandler())
	// run server
	r.Run("0.0.0.0:8888")
}
