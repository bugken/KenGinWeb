package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (communityInfo *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
