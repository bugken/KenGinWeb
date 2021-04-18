package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

// 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db_name"),     // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
