package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	code: 10000,	// 程序中的错误码
	msg: "xxx",		// 提示信息
	data: "xxx" 	// 返回的数据
}
*/

type (
	Response struct {
		Code ResCode     `json:"code"`
		Msg  interface{} `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}
)

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
