package mysql

import (
	"NetClassGinWeb/bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "GinWeb"

// CheckUserExist 检查用户是否存在
func CheckUserExist(userName string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, userName); err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
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

// encryptPassword 加密密码
func encryptPassword(raw string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(raw)))
}

// Login 登录
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id, username, password from user where username = ?"
	if err = db.Get(user, sqlStr, user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotExist
		}
		return
	}

	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrInvalidPassword
	}

	return
}

// GetUserByID 根据ID获取用户信息
func GetUserByID(userID int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username, email, gender, create_time
 				from user where user_id = ?`
	err = db.Select(user, sqlStr, userID)

	return
}
