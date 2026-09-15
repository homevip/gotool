package hlcache

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
)

// 全局 Cache 客户端(外部注入)
var localCache *cache.Cache

// 初始化本地缓存(建议在程序启动时调用)
// 外部项目调用这个方法传入客户端,不再依赖内部 config 路径
func InitLocalCache(client *cache.Cache) {
	localCache = client
}

// LocalCache 本地缓存封装
// args规则:
// S(key)           -> 获取缓存
// S(key, val)      -> 设置永久缓存
// S(key, val, sec) -> 设置缓存，sec秒过期
func S(args ...any) string {
	// 初始化检查(防止未初始化就使用)
	if localCache == nil {
		InitLocalCache(cache.New(0, 5*time.Minute))
	}

	var (
		argN = len(args)
	)

	switch argN {
	case 1:
		key := fmt.Sprint(args[0])
		val, found := localCache.Get(key)
		if found {
			return gconv.String(val)
		}
		return ""

	case 2:
		key := fmt.Sprint(args[0])
		val := gconv.String(args[1])
		// expire=0 永不过期
		localCache.Set(key, val, 0)
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
		localCache.Set(key, val, expire)
		return val

	default:
		// 参数数量非法，返回空，可打错误日志
		return ""
	}
}
