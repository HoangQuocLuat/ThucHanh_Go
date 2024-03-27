package logic

import (
	"context"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
)

type ChatLogic interface {
	Insert(ctx context.Context, chat req.UserChat) (models.RoomChat, error)
}
