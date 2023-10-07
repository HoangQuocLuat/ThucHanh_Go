package controllers

import (
	"log"
	"thuchanh_go/database"
	"thuchanh_go/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

func PostUserRegisLogic(req types.PostUserRegisReq) error {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	logx.Info(req)
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(req.Hashpassword), 14)

	query := "INSERT INTO user_tbl (id, fullname, name, hashpassword, email) VALUES(?, ?, ?, ?, ?)"
	_, err = db.Exec(query, req.ID, req.Fullname, req.Name, hashpassword, req.Email)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
