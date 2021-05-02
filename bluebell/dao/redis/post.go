package redis

import (
	"NetClassGinWeb/bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID, CommunityID int64) (err error) {
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

	// 把帖子ID加入到社区的set里面
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(CommunityID)))
	pipeline.SAdd(cKey, postID)

	// 执行
	_, err = pipeline.Exec()

	return
}

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1

	// 使用RevRange查询
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取IDs
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostTimeZSet)
	}

	return getIDsFromKey(key, p.Page, p.Size)
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

func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// 使用zinterStore 把分区的帖子set与按照帖子分数的zset生成一个新的zset
	// 针对新的zset按照之前的逻辑取数据
	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key减少zinterstore执行的次数
	key := p.Order + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(key).Val() < 1 {
		// 不存在就需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, p.Order)
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据key查询ids
	return getIDsFromKey(key, p.Page, p.Size)
}
