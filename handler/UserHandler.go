package handler

import (
	"net/http"
	"thuchanh_go/controllers"
	"thuchanh_go/types"

	"github.com/gin-gonic/gin"
)

func PostUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.PostUserRegisReq
		var err = ctx.ShouldBindJSON(&req)
		if err != nil {
			panic(err)
		}

		res, err := controllers.PostUserRegisLogic(req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
