package hlcache

import (
	"fmt"
	"testing"
	"time"
)

func TestS(t *testing.T) {

	var (
		cacheKey = "test_key"
		cacheVal = time.Now().Format("2006-01-02 15:04:05")
	)

	cache := S(cacheKey)
	if cache == "" {
		cache = S(cacheKey, cacheVal, "100")
	}

	fmt.Printf("cache: %v\n", cache)
}
