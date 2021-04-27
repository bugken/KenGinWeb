package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/models"
	"NetClassGinWeb/bluebell/thirdparty/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	// 生成post id
	post.ID = snowflake.GenID()

	// 保存到数据库
	return mysql.InsertPost(post)
}

func GetPostByID(postID int64) (data *models.APIPostDetail, err error) {
	// 查询并组合接口需要的数据
	// 查询post信息
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("[GetPostByID]GetPostByID error", zap.Int64("post id", postID), zap.Error(err))
		return data, err
	}

	// 根据作者id查询作者信息
	author, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("[GetPostByID]GetUserByID error", zap.Int64("author id", post.AuthorID), zap.Error(err))
		return data, err
	}

	// 根据社区ID查询社区详情
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("[GetPostByID]GetCommunityDetailByID error",
			zap.Int64("community id", post.CommunityID), zap.Error(err))
		return data, err
	}

	data = &models.APIPostDetail{
		AuthorName:      author.UserName,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子详情列表
func GetPostList(pageSize, pageIndex int64) (data []*models.APIPostDetail, err error) {
	// 查询post信息
	posts, err := mysql.GetPostList(pageSize, pageIndex)
	if err != nil {
		zap.L().Error("[GetPostList]GetPostList error", zap.Error(err))
		return data, err
	}

	data = make([]*models.APIPostDetail, 0, len(posts))
	for _, v := range posts {
		// 根据作者id查询作者信息
		author, err := mysql.GetUserByID(v.AuthorID)
		if err != nil {
			zap.L().Error("[GetPostList]GetUserByID error", zap.Int64("author id", v.AuthorID), zap.Error(err))
			continue
		}

		// 根据社区ID查询社区详情
		community, err := mysql.GetCommunityDetailByID(v.CommunityID)
		if err != nil {
			zap.L().Error("[GetPostList]GetCommunityDetailByID error",
				zap.Int64("community id", v.CommunityID), zap.Error(err))
			continue
		}

		data = append(data, &models.APIPostDetail{
			AuthorName:      author.UserName,
			Post:            v,
			CommunityDetail: community,
		})
	}
	return
}
