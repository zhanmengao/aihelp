package database

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	framework "github.com/zhanmengao/gf/context"
	"time"
)

type THashStore[T IDBProto] struct {
	dbName      string
	keyFormat   string
	keyDesc     string
	fieldFormat string
}

func NewHashStore[T IDBProto](dbName, keyFormat, fieldFormat, keyDesc string) *THashStore[T] {
	return &THashStore[T]{
		dbName:      dbName,
		keyFormat:   keyFormat,
		keyDesc:     keyDesc,
		fieldFormat: fieldFormat,
	}
}

func (p *THashStore[T]) Get(ctx context.Context, key, field string, pData *T, opts ...*Options[T]) (ok bool, err error) {
	ok, err = p.get(ctx, key, field, pData, opts...)
	if err != nil {
		return
	}
	return
}

func (p *THashStore[T]) get(ctx context.Context, key, field string, pData *T, opts ...*Options[T]) (ok bool, err error) {
	cacheKey := framework.CacheKeyPrefix + key + field
	//context有，直接返回
	data, ok := ctx.Value(cacheKey).(T)
	if ok {
		*pData = data
		return
	}
	data = *pData

	command := "HGET"
	IncrOperationCount(command, p.keyFormat, p.fieldFormat, "")
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	bs, err := rdb.HGet(ctx, key, field)
	ObserveDurationMS(command, p.keyFormat, p.fieldFormat, time.Since(opStart), err)
	ObserveDataBytes(command, p.keyFormat, p.fieldFormat, len(bs.Bytes()))

	if bs.IsNil() {
		err = nil
		return
	} else if err != nil {
		return
	}
	ok = true
	err = data.Unmarshal(bs.Bytes())
	if err != nil {
		return
	}

	return
}

func (p *THashStore[T]) GetFromSession(ctx context.Context, key, field string, data *T, opts ...*Options[T]) (ok bool, err error) {
	sess, exist := framework.GetSession(ctx)
	if !exist {
		//Session不存在，走DB
		ok, err = p.get(ctx, key, field, data, opts...)
	} else {
		_, ok, err = p.getFromCache(ctx, sess, key, field, data, opts...)
	}
	return
}

func (p *THashStore[T]) getFromCache(ctx context.Context, c ICache, key, field string, data *T, opts ...*Options[T]) (withCache, ok bool, err error) {
	cacheKey := framework.CacheKeyPrefix + key + field
	obj, ok := c.GetFromCache(ctx, cacheKey)
	if ok {
		var cache T
		if cache, withCache = obj.(T); withCache {
			*data = cache
			return
		}
	}
	ok, err = p.get(ctx, key, field, data, opts...)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	rdb := g.Redis(p.dbName)
	expired, err2 := rdb.TTL(ctx, key)
	if err2 != nil {
	} else if expired > 0 {
		c.SetToCacheTTL(ctx, cacheKey, data, int(expired))
	} else {
		c.SetToCache(ctx, cacheKey, data)
	}
	return
}

func (p *THashStore[T]) Set(ctx context.Context, key, field string, data T) (err error) {
	const command = "SET"
	err = p.set(ctx, key, field, data)
	return
}

func (p *THashStore[T]) set(ctx context.Context, key, field string, data T) (err error) {
	const command = "SET"
	IncrOperationCount(command, p.keyFormat, p.fieldFormat, "")
	bs, err := data.Marshal()
	if err != nil {
		return
	}
	ObserveDataBytes(command, p.keyFormat, p.fieldFormat, len(bs))
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	_, err = rdb.HSet(ctx, key, map[string]interface{}{
		field: bs,
	})
	ObserveDurationMS(command, p.keyFormat, p.fieldFormat, time.Since(opStart), err)
	if err == nil {
		//更新到ctx
		cacheKey := framework.CacheKeyPrefix + key + field
		framework.SetContextValue(ctx, cacheKey, data)
	}
	return
}

func (p *THashStore[T]) SetWithSess(ctx context.Context, key, field string, data T) (err error) {
	var cache bool
	if err = p.set(ctx, key, field, data); err != nil {
		return
	}
	sess, cache := framework.GetSession(ctx)
	if cache {
		cacheKey := framework.CacheKeyPrefix + key + field
		sess.SetToCache(ctx, cacheKey, data)
	}
	return
}

func (p *THashStore[T]) Delete(ctx context.Context, key, field string) (ok bool, err error) {
	const command = "HDEL"
	IncrOperationCount(command, p.keyFormat, field, "")
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	cnt, err := rdb.HDel(ctx, key)
	ok = err == nil && cnt > 0
	ObserveDurationMS(command, p.dbName, field, time.Since(opStart), err)
	return
}
