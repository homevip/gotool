package hlcache

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {

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
