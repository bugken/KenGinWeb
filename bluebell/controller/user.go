package controller

import (
	"NetClassGinWeb/bluebell/logic"
	"NetClassGinWeb/bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[SignUpHandler]请求参数错误", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 2.处理业务逻辑
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("[SignUpHandler]注册失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[LoginHandler]请求参数错误", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("[LoginHandler]登录失败", zap.String("username", p.UserName), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败",
		})
		return
	}

	//返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
