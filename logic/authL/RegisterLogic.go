package authL

import (
	"thuchanh_go/config/db"

	"github.com/zeromicro/go-zero/core/logx"
)

func RegisterLogic(req *RegisterReq) (*Result, error) {

	defer db.CloseDb()

	_, err := Register.Insert(req)
	if err != nil {
		logx.Info(err)
	}

	return &Result{
		Code:    200,
		Message: "đăng ký thành công",
	}, nil
}
