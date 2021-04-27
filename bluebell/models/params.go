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

	// 投票数据
	VoteData struct {
		// UserID 从当前登录用户中获取
		// 帖子ID
		PostID string `json:"post_id,string" binding:"required"`
		// 投票类型 赞成票1 反对票-1 q取消投票
		Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"`
	}
)
