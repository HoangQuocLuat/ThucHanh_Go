package chathandler

import (
	"fmt"
	"log"
	"net/http"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
	"thuchanh_go/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Chat logic.ChatLogic
	Room *ws.Room
}

func NewHandler(c logic.ChatLogic, r *ws.Room) *Handler {
	return &Handler{
		Chat: c,
		Room: r,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req req.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Chat.Insert(&gin.Context{}, req)
	if err != nil {
		c.JSON(http.StatusConflict, models.Response{
			StatusCode: http.StatusConflict,
			Message:    "không tạo được phòng",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Tạo phòng thành công",
		Data:       res,
	})
}

func (h *Handler) GetRooms(c *gin.Context) {
	res, err := h.Chat.Select(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy thông tin thành công",
		Data:       res,
	})
}

// xử lý websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // thử trên postman
		//origin := r.Header.Get("Origin")
		//return origin == "http://localhost:3000"
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	//nâng cấp kết nối HTTP lên kết nối Websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//lấy dữ liệu từ URL
	roomID := c.Param("roomId")
	userID := c.Query("userId")
	username := c.Query("username")

	cl := ws.NewClient(conn, userID, roomID, username, h.Chat)
	ms := &req.Message{
		Content: "A new user has joined the room",
		RoomID:  roomID,
		Sender:  username,
	}

	//Đăng ký một client mới thông qua channal đăng ký
	h.Room.Register <- cl
	// và phát tin nhắn đó
	h.Room.Broadcast <- ms
	//writeMess()
	go cl.WriteMess()
	//readMess()
	cl.ReadMess(h.Room)
	fmt.Print("bbbb", ms)
}

func (h *Handler) OutRoom(c *gin.Context) {
	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	roomID := c.Param("roomId")
	// username := c.Query("username")
	var ReqOutRoom req.ReqOutRoom
	if err := c.ShouldBindJSON(&ReqOutRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print("truy cập thành công outRoom")

	ms := &req.Message{
		Content: "A user exit the room",
		RoomID:  roomID,
		Sender:  "System",
	}

	//phát tin nhắn trong phòng
	h.Room.Broadcast <- ms
	exitClient := &ws.Client{ID: ReqOutRoom.UserID}
	//thoát khỏi phòng
	h.Room.Unregister <- exitClient
}
