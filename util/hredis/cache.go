package hredis

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端(外部注入)
var redisClient *redis.Client

// InitRedis 初始化 Redis 客户端
// 外部项目调用这个方法传入客户端,不再依赖内部 config 路径
func InitRedis(client *redis.Client) {
	redisClient = client
}

// 缓存函数
// args规则:
// S(key)          -> 获取缓存
// S(key, val)     -> 设置永久缓存
// S(key, val, sec)-> 设置缓存，sec秒过期
func S(args ...any) string {
	if redisClient == nil {
		return ""
	}

	var (
		ctx  = context.Background()
		argN = len(args)
	)

	switch argN {
	case 1:
		key := fmt.Sprint(args[0])
		val, err := redisClient.Get(ctx, key).Result()
		if err != nil {
			// 可选:日志打印err,比如gf glog.Error(err)
			return ""
		}
		return val

	case 2:
		key := fmt.Sprint(args[0])
		val := gconv.String(args[1])
		// expire=0 永久不过期
		_, err := redisClient.Set(ctx, key, val, 0).Result()
		if err != nil {
			// log err
		}
		return val

	case 3:
		key := fmt.Sprint(args[0])
		val := gconv.String(args[1])
		sec := gconv.Int(args[2])
		if sec < 0 {
			// 负数非法,直接返回原值,不写入缓存,可加日志
			return val
		}
		expire := time.Duration(sec) * time.Second
		_, err := redisClient.Set(ctx, key, val, expire).Result()
		if err != nil {
			// log err
		}
		return val
	default:
		return ""
	}
}
