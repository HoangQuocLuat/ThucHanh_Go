package types

type GetUserReq struct {
	UserID int64 `uri:"user_id"`
	// Name 	string 	`form:"name"`
	// Password string `json:"password"`
}

type PostUserRegisReq struct {
	ID 				int64   `json:"id"`
	Fullname		string	`json:"fullname"`
	Name 			string  `json:"name"`
	Hashpassword 	string 	`json:"hashpassword"`
	Email 			string 	`json:"email"`
}

type GetUserRes struct {
	ID           	int64  	`json:"id"`
	Fullname     	string 	`json:"fullname"`
	Name         	string 	`json:"name"`
	Hashpassword 	string 	`json:"hashpassword"`
	Email        	string 	`json:"email"`
}