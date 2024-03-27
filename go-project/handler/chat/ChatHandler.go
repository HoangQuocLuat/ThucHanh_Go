package handler

import (
	"log"
	"thuchanh_go/banana"
	"thuchanh_go/logic"
	"thuchanh_go/types/req"

	"github.com/gin-gonic/gin"
)

type UserChatHandler struct {
	UserChat logic.ChatLogic
}

func (c *UserChatHandler) ChatHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws, err := banana.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if err != nil {
			log.Fatal(err)
		}
		defer ws.Close()

		for {
			var msg req.UserChat
			err := ws.ReadJSON(&msg)
			if err != nil {
				log.Printf("error: %v", err)
				break
			}

			err = c.UserChat.Insert(msg)
			if err != nil {
				log.Println("Failed to save message: ", err)
				break
			}
		}
	}
}
