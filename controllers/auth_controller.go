package controllers

import (
	"log"
	"thuchanh_go/database"
	"thuchanh_go/types"
	"thuchanh_go/models"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

func GetUserLogic(req types.GetUserReq) (*types.GetUserRes, error) {
	//connect db mysql
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	logx.Info(req)

	//get data
	var user models.UserTbl

	err = db.QueryRow("SELECT hashpassword FROM user_tbl WHERE name = ?", req.Name).
		Scan(&user.Hashpassword)
	if err != nil {

		logx.Error(err)
		return &types.GetUserRes{
			Result: types.Result{
				Code : 400,
				Message: "sai mật khẩu",
			},
		}, nil
	}
	logx.Info(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Hashpassword), []byte(req.Hashpassword))
	if err != nil {

		logx.Error(err)
		return &types.GetUserRes{
			Result: types.Result{
				Code : 400,
				Message: "sai mật khẩu",
			},
		}, nil
	}
	//xu ly data

	return &types.GetUserRes{
		ID:           user.ID,
		Fullname:     user.Fullname,
		Name:         user.Name,
		Hashpassword: user.Hashpassword,
		Email:        user.Email,
		Result:       types.Result{Code: 200, Message: "thành công"},
	}, nil
}
