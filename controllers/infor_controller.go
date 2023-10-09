package controllers

// import (
// 	"log"
// 	"strconv"
// 	"thuchanh_go/database"
// 	"thuchanh_go/models"
// 	"thuchanh_go/types"

// 	"github.com/zeromicro/go-zero/core/logx"
// )

// func GetUserInfor(req types.GetIdUserReq) (*types.GetUserRes, error) {
// 	db, err := database.DBConn()
// 	defer db.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	logx.Info(req)
// 	var user models.UserTbl
// 	var query = "SELECT id, fullname, email FROM user_tbl WHERE id = "
// 	var row = db.QueryRow(query+ strconv.FormatInt(req.UserID,10))
// 	if err != nil {
// 		logx.Info(err)
// 	}

// 	err = row.Scan(&user.ID, &user.Fullname, &user.Email)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var res = types.GetUserRes(user)

	
// }