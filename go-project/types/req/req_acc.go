package req

type UserRegisReq struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInforReq struct {
	Id string `uri:"userId"`
}
