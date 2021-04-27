package mysql

import (
	"NetClassGinWeb/bluebell/models"
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
		from post limit ? ?`

	posts = make([]*models.Post, 0, 10)
	err = db.Select(&posts, sqlStr, (pageIndex-1)*pageSize, pageSize)

	return
}
