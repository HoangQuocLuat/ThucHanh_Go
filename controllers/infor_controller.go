package controllers

import (
	"log"
	"thuchanh_go/database"
	"thuchanh_go/models"
	"thuchanh_go/types"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetUserInfor(req types.GetIdUserReq) (*types.GetInfor, error) {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	logx.Info(req)
	var user models.UserTbl
	var query = "SELECT id, fullname, email FROM user_tbl WHERE id = ?"
	var row = db.QueryRow(query, req.UserID)
	
	err = row.Scan(&user.ID, &user.Fullname, &user.Email)

	if err != nil {
		logx.Info(err)
		return nil, err
	}

	var res = types.GetInfor{
		ID : user.ID,
		Fullname: user.Fullname,
		Email: user.Email,
	}

	return &res, nil
	
}