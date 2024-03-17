// Code generated by dbtool, DO NOT EDIT.

package gendb

import (
	"context"

	gengocache "github.com/zhanmengao/aihelp/manifest/protobuf/gendb/internal/gengocache"
	genredis "github.com/zhanmengao/aihelp/manifest/protobuf/gendb/internal/genredis"

	basedb "github.com/zhanmengao/aihelp/database"
	pb "github.com/zhanmengao/aihelp/manifest/protobuf/pb"
)

type IPikaDB interface {
	// GetDBChatSession loads DBChatSession from `Pika`
	//   data *pb.DBChatSession if every thing goes well.
	//   ok   bool is false if there is no data.
	//   err  error is not nil if there is something wrong except no data.
	//
	// prometheus metric: QPS, KeySize and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	GetDBChatSession(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBChatSession]) (data *pb.DBChatSession, ok bool, err error)
	// GetDBChatSessionWithCache try get data from cache provider
	// first, if no data in cache, it will fall back to GetDBChatSession
	//   c cache.ICache cache provider
	//
	GetDBChatSessionWithSession(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBChatSession]) (data *pb.DBChatSession, ok bool, err error)
	// SetDBChatSession saves DBChatSession to `Pika`
	//   err error any error when put data to database
	//
	// prometheus metric: QPS, KeySize and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	SetDBChatSession(ctx context.Context, data *pb.DBChatSession) (err error)
	// SetDBChatSessionWithCache try SetDBChatSession first. If success, put it to cache
	//   c cache.ICache cache provider
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	SetDBChatSessionWithSess(ctx context.Context, data *pb.DBChatSession) (err error)
	// GetKeyDBChatSession returns key of DBChatSession
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	GetKeyDBChatSession(UID string) string
	// GetFieldDBChatSession returns field of DBChatSession
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	GetFieldDBChatSession(SessionKey string) string
	// DeleteDBChatSession delete DBChatSession from `Pika`
	//   err  error any error when put data to database
	//
	// prometheus metric: QPS and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
	DeleteDBChatSession(ctx context.Context, UID string, SessionKey string) (ok bool, err error)
	// GetDBUserMessage loads DBUserMessage from `Pika`
	//   data *pb.DBUserMessage if every thing goes well.
	//   ok   bool is false if there is no data.
	//   err  error is not nil if there is something wrong except no data.
	//
	// prometheus metric: QPS, KeySize and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	GetDBUserMessage(ctx context.Context, UID string, opts ...*basedb.Options[*pb.DBUserMessage]) (data *pb.DBUserMessage, ok bool, err error)
	// GetDBUserMessageWithCache try get data from cache provider
	// first, if no data in cache, it will fall back to GetDBUserMessage
	//   c cache.ICache cache provider
	//
	GetDBUserMessageWithSession(ctx context.Context, UID string, opts ...*basedb.Options[*pb.DBUserMessage]) (data *pb.DBUserMessage, ok bool, err error)
	// SetDBUserMessage saves DBUserMessage to `Pika`
	//   err error any error when put data to database
	//
	// prometheus metric: QPS, KeySize and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	SetDBUserMessage(ctx context.Context, data *pb.DBUserMessage) (err error)
	// SetDBUserMessageWithCache try SetDBUserMessage first. If success, put it to cache
	//   c cache.ICache cache provider
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	SetDBUserMessageWithSess(ctx context.Context, data *pb.DBUserMessage) (err error)
	// GetKeyDBUserMessage returns key of DBUserMessage
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	GetKeyDBUserMessage(UID string) string
	// GetFieldDBUserMessage returns field of DBUserMessage
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	GetFieldDBUserMessage() string
	// DeleteDBUserMessage delete DBUserMessage from `Pika`
	//   err  error any error when put data to database
	//
	// prometheus metric: QPS and OperationProcessTime
	//
	// generated from "// @db:hash|Pika|DBUser:%s,UID|DBUserMessage"
	DeleteDBUserMessage(ctx context.Context, UID string) (ok bool, err error)
}

// assert *genredis.Pika{} has implemented IPikaDB
var _ IPikaDB = (*genredis.Pika)(nil)

// NewPikaDBRedis gets a new IPikaDB instance
func NewPikaDBRedis() IPikaDB {
	return &genredis.Pika{}
}

// NewPikaDBCache gets a new IPikaDB instance
func NewPikaDBCache() IPikaDB {
	return &gengocache.Pika{}
}
