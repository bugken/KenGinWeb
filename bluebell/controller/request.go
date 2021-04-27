package controller

import (
	"errors"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户没有登录")

const (
	ContextUserIDKey   = "userID"
	ContextUserNameKey = "userName"
)

// GetCurrentUserID 获取当期用户的用户ID
func GetCurrentUserID(c *gin.Context) (userId int64, err error) {
	v, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok := v.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return userID, nil
}

// GetPageInfo 获取页码信息
func GetPageInfo(c *gin.Context) (int64, int64) {
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	var (
		pageSize  int64
		pageIndex int64
		err       error
	)
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		zap.L().Error("[GetPageInfo]ParseInt error", zap.String("page size", pageIndexStr), zap.Error(err))
		pageSize = 10
	}
	pageIndex, err = strconv.ParseInt(pageIndexStr, 10, 64)
	if err != nil {
		zap.L().Error("[GetPageInfo]ParseInt error", zap.String("page index", pageIndexStr), zap.Error(err))
		pageIndex = 1
	}

	return pageSize, pageIndex
}
