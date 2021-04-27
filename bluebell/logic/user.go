package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/models"
	"NetClassGinWeb/bluebell/thirdparty/jwt"
	"NetClassGinWeb/bluebell/thirdparty/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存在不存在
	if err = mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}

	// 2.生成UID
	userID := snowflake.GenID()

	// 3.构造user对象
	u := &models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}

	// 4.保存到数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}

	// 生成JWT Token
	user.Token, err = jwt.GenToken(user.UserName, user.UserID)

	return
}
