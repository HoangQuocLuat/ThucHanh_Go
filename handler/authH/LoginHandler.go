package authH

import (
	"net/http"
	"thuchanh_go/logic/authL"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req authL.LoginReq
		var err = ctx.ShouldBindJSON(&req)
		if err != nil {
			panic(err)
		}

		//xử lý dữ liệu và đưa ra kết quả
		res, err := authL.LoginLogic(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, res)

	}
}
