package models

type UserTbl struct {
	Username string `db:"username"`
	Fullname string `db:"fullname"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
