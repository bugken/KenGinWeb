package redis

import (
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID int64) (err error) {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 执行
	_, err = pipeline.Exec()

	return
}
