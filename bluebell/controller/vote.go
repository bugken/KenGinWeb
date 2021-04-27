package controller

import (
	"NetClassGinWeb/bluebell/logic"
	"NetClassGinWeb/bluebell/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 投票处理函数
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.VoteData)
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			errData := removeTopStruct(errs.Translate(trans))
			ResponseErrorWithMsg(c, CodeInvalidParam, errData)
			return
		}
	}

	// 获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err = logic.VoteForPost(userID, p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}
