package handler

import (
	"net/http"
	"thuchanh_go/controllers"
	"thuchanh_go/types"

	"github.com/gin-gonic/gin"
)
func GetUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//lay data input
		var req types.GetUserReq
		var err = ctx.ShouldBindUri(&req)
		if err != nil {
			panic(err)
		}

		res, err := controllers.GetUserLogic(req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}