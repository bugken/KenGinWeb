package redis

import (
	"NetClassGinWeb/bluebell/models"
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

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取IDs
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostTimeZSet)
	}

	// 确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	// 使用RevRange查询
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据IDs查询每篇帖子的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	count := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, count)
	//}

	// 使用pipeline一次发送多条命令减少RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return
	}

	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}
