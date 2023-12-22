package models

type UserTbl struct {
	ID           	int64  	`db:"id"`
	Fullname     	string 	`db:"fullname"`
	Name         	string 	`db:"name"`
	Hashpassword 	string 	`db:"hashpassword"`
	Email        	string 	`db:"email"`
}