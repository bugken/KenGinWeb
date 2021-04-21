package models

// 定义请求的参数结构体
type (
	// 注册请求参数
	ParamSignUp struct {
		UserName   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	}

	// 登录请求参数
	ParamLogin struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)
