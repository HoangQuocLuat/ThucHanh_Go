package req

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Message struct {
	RoomID  string `json:"roomid"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type ReqOutRoom struct {
	UserID string `json:"userId"`
}