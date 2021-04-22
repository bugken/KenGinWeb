package controller

import (
	"NetClassGinWeb/bluebell/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ---社区相关文件---

// CommunityHandler 处理Community函数
// TODO:分页显示数据 pageSize pageNum
func CommunityHandler(c *gin.Context) {
	// 获取社区列表信息,以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("[CommunityHandler]GetCommunityList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
	return
}
