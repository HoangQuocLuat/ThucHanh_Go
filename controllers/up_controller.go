package controllers

import (
	"log"
	"strconv"
	"thuchanh_go/database"
	"thuchanh_go/types"

	"github.com/zeromicro/go-zero/core/logx"
)

func PutUserLogic(req types.PutUserReq, reqId types.GetUserReq) error {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	logx.Info(req, reqId)

	up := "UPDATE user_tbl SET hashpassword = ?, email = ? WHERE id = ?"
	_, err = db.Exec(up, req.Hashpassword, req.Email, strconv.FormatInt(reqId.UserID, 10))

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
