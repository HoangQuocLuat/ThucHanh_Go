package controllers

import (
	"log"
	"strconv"
	"thuchanh_go/database"
	"thuchanh_go/models"
	"thuchanh_go/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	var row = db.QueryRow("SELECT * FROM user_tbl WHERE id = " + strconv.FormatInt(req.UserID, 10))
	if err != nil {
		log.Fatal(err)
	}

	err = row.Scan(&user.ID, &user.Fullname, &user.Name, &user.Hashpassword, &user.Email)
	if err != nil {
		return nil, err
	}
	//xu ly data
	var res = types.GetUserRes(user)

	return &res, nil
}