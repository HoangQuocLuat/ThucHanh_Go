package authH

import (
	"net/http"
	"thuchanh_go/logic/authL"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req authL.RegisterReq
		var err = ctx.ShouldBindJSON(&req)
		if err != nil {
			logx.Info(err)
		}

		res, err := authL.RegisterLogic(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
