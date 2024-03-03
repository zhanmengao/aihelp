package database

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	framework "github.com/zhanmengao/gf/context"
	"time"
)

type TStringStore[T IDBProto] struct {
	dbName    string
	keyFormat string
	keyDesc   string
}

func NewStringStore[T IDBProto](dbName, keyFormat, keyDesc string) *TStringStore[T] {
	return &TStringStore[T]{
		dbName:    dbName,
		keyFormat: keyFormat,
		keyDesc:   keyDesc,
	}
}

func (p *TStringStore[T]) Get(ctx context.Context, key string, pData *T, opts ...*Options[T]) (ok bool, err error) {
	ok, err = p.get(ctx, key, pData, opts...)
	if err != nil {
		return
	}
	return
}

func (p *TStringStore[T]) get(ctx context.Context, key string, pData *T, opts ...*Options[T]) (ok bool, err error) {
	cacheKey := framework.CacheKeyPrefix + key
	//context有，直接返回
	data, ok := ctx.Value(cacheKey).(T)
	if ok {
		*pData = data
		return
	}
	data = *pData

	command := "GET"
	IncrOperationCount(command, p.keyFormat, "", "")
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	bs, err := rdb.Get(ctx, key)
	ObserveDurationMS(command, p.keyFormat, "", time.Since(opStart), err)
	ObserveDataBytes(command, p.keyFormat, "", len(bs.Bytes()))

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

func (p *TStringStore[T]) GetFromSession(ctx context.Context, key string, data *T, opts ...*Options[T]) (ok bool, err error) {
	sess, exist := framework.GetSession(ctx)
	if !exist {
		ok, err = p.Get(ctx, key, data, opts...)
	} else {
		_, ok, err = p.getFromCache(ctx, sess, key, data, opts...)
	}
	return
}

func (p *TStringStore[T]) getFromCache(ctx context.Context, c ICache, key string, data *T, opts ...*Options[T]) (withCache, ok bool, err error) {
	cacheKey := framework.CacheKeyPrefix + key
	obj, ok := c.GetFromCache(ctx, cacheKey)
	if ok {
		var cache T
		if cache, withCache = obj.(T); withCache {
			*data = cache
			return
		}
	}
	ok, err = p.get(ctx, key, data, opts...)
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
		c.SetToCacheTTL(ctx, cacheKey, *data, int(expired))
	} else {
		c.SetToCache(ctx, cacheKey, *data)
	}
	return
}

func (p *TStringStore[T]) Set(ctx context.Context, key string, data T) (err error) {
	err = p.set(ctx, key, data)
	return
}

func (p *TStringStore[T]) set(ctx context.Context, key string, data T) (err error) {
	const command = "SET"
	IncrOperationCount(command, p.keyFormat, "", "")
	bs, err := data.Marshal()
	if err != nil {
		return
	}
	ObserveDataBytes(command, p.keyFormat, "", len(bs))
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	_, err = rdb.Set(ctx, key, bs)
	ObserveDurationMS(command, p.keyFormat, "", time.Since(opStart), err)
	if err == nil {
		//更新到ctx
		cacheKey := framework.CacheKeyPrefix + key
		framework.SetContextValue(ctx, cacheKey, data)
	}
	return
}

func (p *TStringStore[T]) SetWithSess(ctx context.Context, key string, data T) (err error) {
	var isCache bool
	err = p.set(ctx, key, data)
	if err != nil {
		return
	}
	sess, isCache := framework.GetSession(ctx)
	if isCache {
		cacheKey := framework.CacheKeyPrefix + key
		sess.SetToCache(ctx, cacheKey, data)
	}
	return
}

func (p *TStringStore[T]) SetEX(ctx context.Context, key string, data T, ttl int) (err error) {
	err = p.setEX(ctx, key, data, ttl)
	return
}

func (p *TStringStore[T]) setEX(ctx context.Context, key string, data T, ttl int) (err error) {
	const command = "SETEX"
	IncrOperationCount(command, p.keyFormat, "", "")
	bs, err := data.Marshal()
	if err != nil {
		return
	}
	ObserveDataBytes(command, p.keyFormat, "", len(bs))
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	err = rdb.SetEX(ctx, key, bs, int64(ttl))
	ObserveDurationMS(command, p.keyFormat, "", time.Since(opStart), err)
	if err == nil {
		//更新到ctx
		cacheKey := framework.CacheKeyPrefix + key
		framework.SetContextValue(ctx, cacheKey, data)
	}
	return
}

func (p *TStringStore[T]) SetEXWithSession(ctx context.Context, key string, data T, ttl int) (err error) {
	var isCache bool
	err = p.setEX(ctx, key, data, ttl)
	if err != nil {
		return
	}
	sess, isCache := framework.GetSession(ctx)
	if isCache {
		cacheKey := framework.CacheKeyPrefix + key
		sess.SetToCacheTTL(ctx, cacheKey, data, ttl)
	}
	return
}

func (p *TStringStore[T]) Delete(ctx context.Context, key string) (ok bool, err error) {
	const command = "DEL"
	IncrOperationCount(command, p.keyFormat, "", "")
	rdb := g.Redis(p.dbName)
	opStart := time.Now()
	cnt, err := rdb.Del(ctx, key)
	ok = err == nil && cnt > 0
	ObserveDurationMS(command, p.dbName, "", time.Since(opStart), err)
	return
}
