package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)

var (
	rdb *redis.Client
)

// 初始化连接
func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func DemoRedisV8() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		fmt.Printf("init redis client failed, error:%v\n", err.Error())
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Printf("redis set failed, error:%v\n", err.Error())
		return
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Printf("redis get key failed, error:%v\n", err.Error())
		return
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Printf("redis get key2 failed, error:%v\n", err.Error())
		return
	} else {
		fmt.Println("key2", val2)
	}
}

/*
Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
	这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间(RTT)
*/
func PipelineDemo() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		fmt.Printf("init redis client failed, error:%v\n", err.Error())
		return
	}

	pipe := rdb.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)
}

func PipelineDemo2() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		fmt.Printf("init redis client failed, error:%v\n", err.Error())
		return
	}

	var incr *redis.IntCmd
	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr.Val(), err)
}

/*
Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
TxPipeline总体上类似于上面的Pipeline，但是它内部会使用MULTI/EXEC包裹排队的命令。
*/
func DemoTxPipeline() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		fmt.Printf("init redis client failed, error:%v\n", err.Error())
		return
	}

	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)
}

func DemoWatch() {
	ctx := context.Background()
	if err := InitClient(); err != nil {
		fmt.Printf("init redis client failed, error:%v\n", err.Error())
		return
	}

	// 监视watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("watch exec failed, error:%v\n", err.Error())
		return
	}
	fmt.Printf("watch exec success.\n")
}
