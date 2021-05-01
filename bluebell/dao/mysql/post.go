package mysql

import (
	"NetClassGinWeb/bluebell/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func InsertPost(post *models.Post) (err error) {
	sqlStr := "insert into post(post_id, author_id, community_id, title, content)" +
		"values(?,?,?,?,?)"

	_, err = db.Exec(sqlStr, post.ID, post.AuthorID, post.CommunityID, post.Title, post.Content)
	return
}

func GetPostByID(postID int64) (post *models.Post, err error) {
	sqlStr := `select post_id, author_id, community_id, title, content,
		create_time from post where post_id = ?`

	post = new(models.Post)
	err = db.Select(post, sqlStr, postID)

	return
}

func GetPostList(pageSize, pageIndex int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, author_id, community_id, title, content,	create_time 
		from post order by create_time desc limit ? ?`

	posts = make([]*models.Post, 0, 10)
	err = db.Select(&posts, sqlStr, (pageIndex-1)*pageSize, pageSize)

	return
}

// 根据给定的ids列表去查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
		from post where post_id in (?) order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&postList, query, args...) // !!!

	return
}
