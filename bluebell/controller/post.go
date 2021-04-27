package controller

import (
	"NetClassGinWeb/bluebell/logic"
	"NetClassGinWeb/bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数以及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
	return
}

// GetPostDetailHandler 获取帖子详情函数
func GetPostDetailHandler(c *gin.Context) {
	// 获取帖子id
	strID := c.Param("id")
	postID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据id取出帖子数
	data, err := logic.GetPostByID(postID)
	if err != nil {
		zap.L().Error("[GetPostDetailHandler]GetPostDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)
	return
}
