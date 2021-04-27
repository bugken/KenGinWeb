package controller

import (
	"errors"

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
