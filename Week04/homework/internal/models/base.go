package models

import redis "github.com/wlxpkg/base/cache"

// cache 实例化redis cache
var cache = redis.NewCache()

// ttl 默认缓存时长
var ttl = 60 * 60 * 24 * 365

// Cacher 缓存接口
type Cacher interface {
	getCacheKey() string
	getCacheTTL() int
}

// cacheKey 获取缓存 key
func cacheKey(c Cacher) string {
	return c.getCacheKey()
}

// cacheTTL 获取缓存时长
func cacheTTL(c Cacher) int {
	return c.getCacheTTL()
}
