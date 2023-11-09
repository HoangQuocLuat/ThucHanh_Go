package handler

// import (
// 	"net/http"
// 	"thuchanh_go/controllers"
// 	"thuchanh_go/types"

// 	"github.com/gin-gonic/gin"
// 	"github.com/zeromicro/go-zero/core/logx"
// )

// func GetUserHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var req types.GetIdUserReq
// 		var err = ctx.ShouldBindUri(&req)
// 		if err != nil {
// 			logx.Info(err)
// 		}
		
// 		res, err := controllers.GetUserInfor(req)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, nil)
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, res)
// 	}
// }