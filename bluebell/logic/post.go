package logic

import (
	"NetClassGinWeb/bluebell/dao/mysql"
	"NetClassGinWeb/bluebell/dao/redis"
	"NetClassGinWeb/bluebell/models"
	"NetClassGinWeb/bluebell/thirdparty/snowflake"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	// 生成post id
	post.ID = snowflake.GenID()

	// 保存到数据库
	if err = mysql.InsertPost(post); err != nil {
		return
	}

	// 更新redis
	if err = redis.CreatePost(post.ID, post.CommunityID); err != nil {
		return
	}

	return
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

// GetPostList2 获取帖子详情列表
func GetPostList2(p *models.ParamPostList) (data []*models.APIPostDetail, err error) {
	// 根据参数去redis获取帖子id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("GetPostIDsInOrder success but get 0 data")
		return
	}

	// 根据id去数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查询好每篇帖子的投票数据
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子详细信息查询出来
	for idx, v := range posts {
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
			VoteNum:         voteData[idx],
			Post:            v,
			CommunityDetail: community,
		})
	}
	return
}

func GetPostCommunityList(p *models.ParamPostList) (data []*models.APIPostDetail, err error) {
	// 根据参数去redis获取帖子id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("GetCommunityPostIDsInOrder success but get 0 data")
		return
	}

	// 根据id去数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查询好每篇帖子的投票数据
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子详细信息查询出来
	for idx, v := range posts {
		// 根据作者id查询作者信息
		author, err := mysql.GetUserByID(v.AuthorID)
		if err != nil {
			zap.L().Error("[GetPostList]GetUserByID error", zap.Int64("author id", v.AuthorID), zap.Error(err))
			continue
		}

		// 根据社区ID查询社区详情
		community, err := mysql.GetCommunityDetailByID(v.CommunityID)
		if err != nil {
			zap.L().Error("[GetPostCommunityList]GetCommunityDetailByID error",
				zap.Int64("community id", v.CommunityID), zap.Error(err))
			continue
		}

		data = append(data, &models.APIPostDetail{
			AuthorName:      author.UserName,
			VoteNum:         voteData[idx],
			Post:            v,
			CommunityDetail: community,
		})
	}
	return
}

// GetPostListNew  将两个查询帖子列表逻辑合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data []*models.APIPostDetail, err error) {
	// 根据请求参数的不同，执行不同的逻辑。
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetPostCommunityList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err),
			zap.Int64("CommunityID", p.CommunityID))
		return nil, err
	}
	return
}
