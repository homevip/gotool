package hredis

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {

	// 初始化 Redis（把项目的 redis 客户端传给工具包）
	// hredis.InitRedis(config.GetRedisDB())

	var (
		cacheKey = "test_key"
		cacheVal = "test-data"
	)

	cache := S(cacheKey)
	if cache == "" {
		cache = S(cacheKey, cacheVal, 10)
	}
	fmt.Printf("cache: %v\n", cache)

}
