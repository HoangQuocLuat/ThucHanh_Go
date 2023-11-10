package authL

import (
	"database/sql"
	"thuchanh_go/config/db"
	"thuchanh_go/types"
	"thuchanh_go/utils"
)
type UserRes struct {
	Result Result     `json:"result"`
	Data   User 	  `json:"data"`
}
type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Register interface {
	Insert() (sql.Result, error)
}

type RegisterReq struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterReq) Insert() (sql.Result, error) {
	db.ConnectDb()

	hashpassword := utils.Hashpassword(req.Password)

	query := "INSERT INTO user_tbl (username, fullname, email, password) VALUES(?, ?, ?, ?)"
	res, err := db.DB.Exec(query, req.Username, req.Fullname, req.Email, hashpassword)

	return res, err
}

type Login interface {
	Find() (User, error)
	
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Fullname string `json: "fullname"`
	Email    string `json: "email"`
	Token    Token  `json: "token"`
}
type Token struct {
	Token string `json:"token"`
}

func () Find(req *types.LoginReq)(User, error) {
	db.ConnectDb()
	var user User

	row := db.DB.QueryRow("SELECT username,fullname, email, password  FROM user_tbl WHERE username = ?", req.Username)
	err := row.Scan(&user.Username, &user.Fullname, &user.Email, &user.Password)

	return user, err
}
