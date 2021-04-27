package mysql

import (
	"NetClassGinWeb/bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community;"

	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("[GetCommunityList]no data in db.")
			err = nil
			return
		}
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (communityInfo *models.CommunityDetail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time " +
		"from community where community_id = ?;"

	communityInfo = new(models.CommunityDetail)
	if err = db.Get(communityInfo, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("[GetCommunityDetail]no data in db.")
			err = ErrInvalidID
			return
		}
	}
	return
}
