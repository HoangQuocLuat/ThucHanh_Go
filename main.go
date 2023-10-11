package main

import (
	"thuchanh_go/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	//r.GET("user/infor/:user_id", handler.GetUserHandler())
	r.PUT("user/upacc/:user_id", handler.PutUserHandler())

	

	// run server
	r.Run("127.0.0.1:8888")
}
