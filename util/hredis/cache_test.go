package hredis

import (
	"fmt"
	"testing"
	"time"
)

func TestS(t *testing.T) {

	// 初始化 Redis（把项目的 redis 客户端传给工具包）
	// hredis.InitRedis(config.GetRedisDB())

	var (
		CacheKey = "test_key"
		CacheVal = time.Now().Format("2006-01-02 15:04:05")
	)

	cache := S(CacheKey)
	if cache == "" {
		cache = S(CacheKey, CacheVal, "100")
	}
	fmt.Printf("cache: %v\n", cache)

}
