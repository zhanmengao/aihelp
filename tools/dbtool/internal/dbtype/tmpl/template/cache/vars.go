package cache

import (
	"fmt"
	"forevernine.com/midplat/base_libs/xtime"
	"sync"
)

var (
	cacheMap = sync.Map{}
)

type redisCache struct {
	ExpiredTime int64
	DBData      interface{}
}

func getLRUCache(key, field string) (data interface{}, ok bool) {
	iBody, ok := cacheMap.Load(fmt.Sprintf("%s_%s", key, field))
	if !ok {
		return
	}
	cache, ok := iBody.(*redisCache)
	if !ok {
		return
	}
	//判断是否过期
	if cache.ExpiredTime > 0 && cache.ExpiredTime < xtime.Unix() {
		ok = false
		return
	}
	data = cache.DBData
	return
}

func setLRUCache(key, field string, ttl int64, data interface{}) {
	cacheMap.Store(fmt.Sprintf("%s_%s", key, field), &redisCache{
		ExpiredTime: xtime.Unix() + ttl,
		DBData:      data,
	})
}

func delLRUCache(key, field string) {
	cacheMap.Delete(fmt.Sprintf("%s_%s", key, field))
}
