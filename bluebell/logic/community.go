package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
