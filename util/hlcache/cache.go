package hlcache

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
)

// 全局go-cache实例,对应原来的global.RedisDB
var localCache *cache.Cache

// 初始化本地缓存(建议在程序启动时调用)
func InitLocalCache() {
	// 默认过期时间设为0(永不过期),清理间隔设为5分钟
	localCache = cache.New(0, 5*time.Minute)
}

// 缓存函数
// 参数1: key
// 参数2: val
// 参数3: 缓存时间/秒
func S(any ...any) string {
	var (
		len = len(any)
	)

	// 初始化检查(防止未初始化就使用)
	if localCache == nil {
		InitLocalCache()
	}

	// 获取值:仅传入key一个参数时
	if len == 1 {
		key := fmt.Sprint(any[0])
		val, found := localCache.Get(key)
		if found {
			return gconv.String(val)
		}
		return ""
	}

	// 设置值:传入key、val、expire三个参数时
	if len == 3 {
		var (
			key    = fmt.Sprint(any[0])
			val    = gconv.String(any[1])
			expire = time.Duration(gconv.Int(any[2])) * time.Second
		)

		// go-cache中0秒表示永不过期(和Redis逻辑一致)
		localCache.Set(key, val, expire)
		return val
	}

	// 参数数量错误时返回空字符串
	return ""
}
