package controller

import (
	"NetClassGinWeb/bluebell/logic"
	"NetClassGinWeb/bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[SignUpHandler]请求参数错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": "请求参数错误"})
		return
	}

	// 2.处理业务逻辑
	logic.SignUp(p)

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
