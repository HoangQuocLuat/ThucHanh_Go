package ws

import (
	"context"
	"log"
	"thuchanh_go/logic"
	"thuchanh_go/types/req"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *req.Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	UserName string `json:"username"`
	Chat     logic.ChatLogic
}

func NewClient(conn *websocket.Conn, id, roomID, userName string, c logic.ChatLogic) *Client {
	return &Client{
		Conn:     conn,
		Message:  make(chan *req.Message),
		ID:       id,
		RoomID:   roomID,
		UserName: userName,
		Chat:     c,
	}
}
func (c *Client) WriteMess() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				return
			}
			err := c.Conn.WriteJSON(message)
			if err != nil {
				log.Println("error writing message:", err)
				return
			}

		}
	}
}

func (c *Client) ReadMess(room *Room) {
	defer func() {
		room.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message req.Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
		}
		message.RoomID = c.RoomID
		message.Sender = c.UserName
		err = c.Chat.InsertMess(context.Background(), message)
		if err != nil {
			log.Println("Error inserting message:", err)
		}
		room.Broadcast <- &message
	}
}
