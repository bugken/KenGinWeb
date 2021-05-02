package redis

// redis key 注意使用命名空间的方式，防止冲突以及方便业务区分

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"
	KeyPostScoreZSet       = "post:score"
	KeyPostVotedZSetPrefix = "post:voted" // 记录用户投票类型,参数为帖子post_id
	KeyCommunitySetPrefix  = "community:" // 保存每个分区下的帖子
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
