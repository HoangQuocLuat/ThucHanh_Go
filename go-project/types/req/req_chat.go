package req

type UserChat struct {
	Message    string 	`json:"message"`
	CreateDate string	`json:"createdate"`
}