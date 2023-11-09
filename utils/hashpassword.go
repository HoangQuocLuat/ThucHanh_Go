package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

// func Hashpassword(req string) ([]byte, error) {
// 	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
// 	if err != nil {
// 		logx.Info(err)
// 	}
// 	return hashpassword, nil
// }

func Hashpassword(pass string) ([]byte) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		logx.Info(err)
	}
	return hashpassword
}
