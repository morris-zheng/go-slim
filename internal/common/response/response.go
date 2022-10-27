package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

type FailStruct struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func Fail(ctx *gin.Context, msg string, code int) {
	ctx.JSON(http.StatusBadRequest, FailStruct{
		Msg:  msg,
		Code: code,
	})
}
