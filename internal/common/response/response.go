package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data     any
	Msg      string
	Code     int
	HttpCode int
}

func Success(ctx *gin.Context, resp Response) {
	if resp.HttpCode == 0 {
		resp.HttpCode = http.StatusOK
	}

	ctx.JSON(resp.HttpCode, resp.Data)
}

func Fail(ctx *gin.Context, resp Response) {
	if resp.HttpCode == 0 {
		resp.HttpCode = http.StatusBadRequest
	}

	r := struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}{
		Msg:  resp.Msg,
		Code: resp.Code,
	}
	ctx.JSON(resp.HttpCode, r)
}
