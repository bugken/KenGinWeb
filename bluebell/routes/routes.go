package routes

import (
	"NetClassGinWeb/bluebell/controller"
	"NetClassGinWeb/bluebell/middleware"
	"NetClassGinWeb/webginbase/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.POST("/signup", controller.SignUpHandler)
	// 登陆路由
	r.POST("/login", controller.LoginHandler)
	// ping-pong路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Pong"})
	})

	// 路由分组 /api/v1
	{
		v1 := r.Group("/api/v1")
		// 接下来的路由使用JWT认证中间件
		v1.Use(middleware.JWTAuthMiddleware())

		// 获取community路由
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 帖子相关接口
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)

		// 投票功能
		v1.POST("/vote", controller.PostVoteHandler)
	}

	return r
}
