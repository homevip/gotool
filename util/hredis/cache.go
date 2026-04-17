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
// 外部项目调用这个方法传入客户端，不再依赖内部 config 路径
func InitRedis(client *redis.Client) {
	redisClient = client
}

// 缓存函数
// 参数1: key
// 参数2: val
// 参数3: 缓存时间/秒
func S(any ...any) string {

	// 安全判断
	if redisClient == nil {
		return ""
	}

	var (
		ctx = context.Background()
		len = len(any)
	)

	// 获取值
	if len == 1 {
		var (
			key = fmt.Sprint(any[0])
		)

		val, _ := redisClient.Get(ctx, key).Result()
		return val
	}

	// 设置值
	if len == 3 {
		var (
			key    = fmt.Sprint(any[0])
			val    = gconv.String(any[1])
			expire = time.Duration(gconv.Int(any[2])) * time.Second
		)

		_, _ = redisClient.Set(ctx, key, val, expire).Result() // 0 表示没有过期时间
		return val
	}

	return ""
}
