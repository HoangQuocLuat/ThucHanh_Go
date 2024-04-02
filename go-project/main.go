// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"thuchanh_go/config"
	"thuchanh_go/database"
	handler "thuchanh_go/handler/account"
	chathandler "thuchanh_go/handler/chat"
	logicacc "thuchanh_go/logic/account"
	chatlogic "thuchanh_go/logic/chat"
	"thuchanh_go/redis"
	router "thuchanh_go/router/acc"
	"thuchanh_go/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	// Xử lý tùy chọn dòng lệnh
	configFile := flag.String("f", "config.yaml", "Path to YAML config file")
	flag.Parse()

	// Đọc cấu hình từ tệp YAML
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Kết nối database
	sql := &database.Sql{
		UserName: cfg.Database.Username,
		Password: cfg.Database.Password,
		DbName:   cfg.Database.DbName,
	}
	sql.Connect()
	defer sql.Close()

	// Kết nối Redis
	redis := &redis.Redis{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	}
	redis.Connect()

	//Hand Chat
	room := ws.NewRoom()
	wsHandler := chathandler.Handler{
		Chat: chatlogic.NewChatLogic(sql),
		Room: room,
	}

	// Khởi tạo Handler
	go room.Run()

	//Hand Account
	userHandler := handler.AccountHandler{
		UserLogic: logicacc.NewAccRegisterLogic(sql),
		Rd:        redis,
	}

	//khởi tạo gin
	r := gin.New()
	// Router
	api := router.API{
		Gin:        r,
		AccHandler: userHandler,
		WebHandler: wsHandler,
	}
	api.SetupRoute()

	// Chạy server
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	r.Run(address)
}
