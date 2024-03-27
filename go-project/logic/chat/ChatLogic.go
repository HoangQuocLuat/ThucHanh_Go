package chat

import (
	"context"
	"thuchanh_go/database"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
)

type ChatLogic struct {
	sql *database.Sql
}

func NewChatLogic(sql *database.Sql) logic.ChatLogic {
	return &ChatLogic{
		sql: sql,
	}
}

func (c *ChatLogic) Insert(ctx context.Context, chatreq req.UserChat) (models.RoomChat, error) {
	statement := `INSERT INTO roomchat (id, message, createdate)
	VALUES(:id, message, createdate)`

	_, err := a.sql.Db.NamedExecContext(ctx, statement, chatreq)
}
