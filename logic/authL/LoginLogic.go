package authL

import (
	"thuchanh_go/config/db"
	"thuchanh_go/constt"
	"thuchanh_go/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

func LoginLogic(req *LoginReq) (*UserRes, error) {
	defer db.CloseDb()
	user, err := Login.Find(req)
	if err != nil {
		logx.Info(err)
	}
	token, err := utils.GenerateJWT(user.Username, constt.SecretKey)
	if err != nil {
		logx.Info(err)
	}
	return &UserRes{
		Result: Result{
			Code:    200,
			Message: "thành công",
		},
		Data: User{
			Username: user.Username,
			Password: "",
			Fullname: user.Fullname,
			Email:    user.Email,
			Token: Token{
				Token: token,
			},
		},
	}, nil
}
