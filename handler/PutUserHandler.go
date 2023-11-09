 package handler

// import (
// 	"net/http"
// 	"thuchanh_go/controllers"
// 	"thuchanh_go/types"

// 	"github.com/gin-gonic/gin"
// )

// func PutUserHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		var req types.PutUserReq
// 		var err = ctx.ShouldBindJSON(&req)
// 		if err != nil {
// 			panic(err)
// 		}

// 		var id types.GetIdUserReq
// 		err = ctx.ShouldBindUri(&id)
// 		if err != nil {
// 			panic(err)
// 		}

// 		err = controllers.PutUserLogic(req, id)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, nil)
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{"mes": "Tài khoản đã được cập nhật thành công"})
// 	}
// }
