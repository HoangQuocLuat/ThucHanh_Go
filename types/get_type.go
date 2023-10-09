package types

type GetIdUserReq struct {
	UserID int64 `uri:"user_id"`
	// Name 	string 	`form:"name"`
	// Password string `json:"password"`
}
type GetUserReq struct {
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
}

type PostUserRegisReq struct {
	ID           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
	Email        string `json:"email"`
}

type GetUserRes struct {
	ID           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
	Email        string `json:"email"`
	Result       Result `json:"result"`
	Token		 Token 	`json:"token"`
}

type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Token  string `json:"token"`
}

type PutUserReq struct {
	Hashpassword string `json:"hashpassword"`
	Email        string `json:"email"`
}
