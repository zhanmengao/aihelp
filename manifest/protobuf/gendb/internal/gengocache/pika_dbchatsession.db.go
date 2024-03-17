package cache

import (
	"context"
	"fmt"
	basedb "github.com/zhanmengao/aihelp/database"
	pb "github.com/zhanmengao/aihelp/manifest/protobuf/pb"
)

// GetKeyDBChatSession returns key of DBChatSession
//
// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
func (p *Pika) GetKeyDBChatSession(UID string) string {
	return fmt.Sprintf(keyDBChatSession, UID)
}

// GetFieldDBChatSession returns field of DBChatSession
//
// generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
func (p *Pika) GetFieldDBChatSession(SessionKey string) string {
	return fmt.Sprintf(fieldDBChatSession, SessionKey)
}

func (p *Pika) GetDBChatSession(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBChatSession]) (data *pb.DBChatSession, ok bool, err error) {
	key := p.GetKeyDBChatSession(UID)
	field := p.GetFieldDBChatSession(SessionKey)
	iData, ok := getLRUCache(key, field)
	if ok {
		data, ok = iData.(*pb.DBChatSession)
	} else {
		data = &pb.DBChatSession{}
		data.UID = UID
		data.SessionKey = SessionKey
	}
	return
}

func (p *Pika) GetDBChatSessionWithSession(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBChatSession]) (data *pb.DBChatSession, ok bool, err error) {
	return p.GetDBChatSession(ctx, UID, SessionKey)
}

func (p *Pika) SetDBChatSession(ctx context.Context, data *pb.DBChatSession) (err error) {
	key := p.GetKeyDBChatSession(data.UID)
	field := p.GetFieldDBChatSession(data.SessionKey)
	setLRUCache(key, field, 0, data)
	return
}

func (p *Pika) SetDBChatSessionWithSess(ctx context.Context, data *pb.DBChatSession) (err error) {
	return p.SetDBChatSession(ctx, data)
}

func (p *Pika) SetEXDBChatSession(ctx context.Context, data *pb.DBChatSession, ttl int) (err error) {
	key := p.GetKeyDBChatSession(data.UID)
	field := p.GetFieldDBChatSession(data.SessionKey)
	setLRUCache(key, field, int64(ttl), data)
	return
}

func (p *Pika) SetEXDBChatSessionWithSess(ctx context.Context, data *pb.DBChatSession, ttl int) (err error) {
	return p.SetEXDBChatSession(ctx, data, ttl)
}

// DeleteDBChatSession is generated from "// @db:hash|Pika|DBUser:%s,UID|DBChatSession:%s,SessionKey"
//
//	err  error any error when put data to database
//
// if monitor is enabled on init this operation records prometheus
// metric: QPS and OperationProcessTime
func (p *Pika) DeleteDBChatSession(ctx context.Context, UID string, SessionKey string) (ok bool, err error) {
	key := p.GetKeyDBChatSession(UID)
	field := p.GetFieldDBChatSession(SessionKey)
	delLRUCache(key, field)
	return
}
