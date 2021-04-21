package mysql

import (
	"NetClassGinWeb/bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

var secret = "GinWeb"

// CheckUserExist 检查用户是否存在
func CheckUserExist(userName string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, userName); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}

	return
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := "insert into user(user_id, username, password) values(?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)

	return
}

func encryptPassword(raw string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(raw)))
}
