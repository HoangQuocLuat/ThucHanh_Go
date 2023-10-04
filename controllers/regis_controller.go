package controllers

import (
	"log"
	"thuchanh_go/database"
	"thuchanh_go/types"

	"github.com/zeromicro/go-zero/core/logx"
)
func PostUserRegisLogic(req types.PostUserRegisReq) (error) {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	logx.Info(req)
	query := "INSERT INTO user_tbl (id, fullname, name, hashpassword, email) VALUES(?, ?, ?, ?, ?)"
	_, err = db.Exec(query, req.ID, req.Fullname, req.Name, req.Hashpassword, req.Email)
	
	if err != nil {
		log.Fatal(err)
	}

	return nil
}