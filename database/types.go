package database

import (
	"context"
	"github.com/gogo/protobuf/proto"
)

type IDBProto interface {
	proto.Marshaler
	proto.Unmarshaler
	proto.Sizer
	String() string
	GetDBKey() string
	GetDBKeyFormat() string
	GetDBField() string
	GetDBFieldFormat() string
	GetDBKeyName() string
	GetDBName() string
	CopyFrom(other IDBProto) (err error)
}

type ICache interface {
	GetFromCache(ctx context.Context, key string) (data interface{}, ok bool)
	SetToCache(ctx context.Context, key string, val interface{})
	SetToCacheTTL(ctx context.Context, key string, val interface{}, ttl int)
}

type Options[T IDBProto] struct {
	AutoSetCall TSetDBFunc[T] //回写的回调，Set可以直接填入，SetEx传入委托函数
}

type TSetDBFunc[T IDBProto] func(ctx context.Context, data T) (err error)

func NewOptions[T IDBProto]() *Options[T] {
	return &Options[T]{}
}

func (o *Options[T]) WithAutoSetCall(cb TSetDBFunc[T]) *Options[T] {
	o.AutoSetCall = cb
	return o
}
