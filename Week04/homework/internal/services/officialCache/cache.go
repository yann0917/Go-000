/*
 * @Author: qiuling
 * @Date: 2019-09-24 18:46:31
 * @Last Modified by: yubb
 * @Last Modified time: 2019-12-19 14:26:36
 */
package officialCache

import (
	"time"

	redis "github.com/wlxpkg/base/cache"
)

type Cache struct {
	client *redis.Cache
}

func NewCache() *Cache {
	return &Cache{
		client: redis.NewCache(),
	}
}

func (c *Cache) Get(key string) interface{} {
	var result interface{}
	err := c.client.Get(key, &result)
	if err != nil {
		return nil
	}
	return result
}

// IsExist check value exists in redis.
func (c *Cache) IsExist(key string) bool {
	return c.client.Exists(key)
}

//Set cached value with key and expire time.
func (c *Cache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	time := int(timeout / time.Second)
	c.client.Set(key, val, time)
	return nil
}

//Delete delete value in redis.
func (c *Cache) Delete(key string) error {
	c.client.Del(key)
	return nil
}
