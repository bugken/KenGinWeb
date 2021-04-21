package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/models"
	"NetClassGinWeb/bluebell/thirdparty/snowflake"
)

// 存放业务逻辑的代码

func SignUp(param *models.ParamSignUp) {
	// 1.判断用户存在不存在
	mysql.QueryUserByUserName()

	// 2.生成UID
	snowflake.GenID()

	// 3.保存到数据库
	mysql.InsertUser()

}
