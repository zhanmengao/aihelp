package cache

import (
	"context"
	"fmt"
	basedb "github.com/zhanmengao/aihelp/database"
	pb "github.com/zhanmengao/aihelp/manifest/protobuf/pb"
)

// GetKeyDBUserMessage returns key of DBUserMessage
//
// generated from "//@db:hash|Pika|DBUserMessage:%s,UID|Session:%s,SessionKey"
func (p *Pika) GetKeyDBUserMessage(UID string) string {
	return fmt.Sprintf(keyDBUserMessage, UID)
}

// GetFieldDBUserMessage returns field of DBUserMessage
//
// generated from "//@db:hash|Pika|DBUserMessage:%s,UID|Session:%s,SessionKey"
func (p *Pika) GetFieldDBUserMessage(SessionKey string) string {
	return fmt.Sprintf(fieldDBUserMessage, SessionKey)
}

func (p *Pika) GetDBUserMessage(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBUserMessage]) (data *pb.DBUserMessage, ok bool, err error) {
	key := p.GetKeyDBUserMessage(UID)
	field := p.GetFieldDBUserMessage(SessionKey)
	iData, ok := getLRUCache(key, field)
	if ok {
		data, ok = iData.(*pb.DBUserMessage)
	} else {
		data = &pb.DBUserMessage{}
		data.UID = UID
		data.SessionKey = SessionKey
	}
	return
}

func (p *Pika) GetDBUserMessageWithSession(ctx context.Context, UID string, SessionKey string, opts ...*basedb.Options[*pb.DBUserMessage]) (data *pb.DBUserMessage, ok bool, err error) {
	return p.GetDBUserMessage(ctx, UID, SessionKey)
}

func (p *Pika) SetDBUserMessage(ctx context.Context, data *pb.DBUserMessage) (err error) {
	key := p.GetKeyDBUserMessage(data.UID)
	field := p.GetFieldDBUserMessage(data.SessionKey)
	setLRUCache(key, field, 0, data)
	return
}

func (p *Pika) SetDBUserMessageWithSess(ctx context.Context, data *pb.DBUserMessage) (err error) {
	return p.SetDBUserMessage(ctx, data)
}

func (p *Pika) SetEXDBUserMessage(ctx context.Context, data *pb.DBUserMessage, ttl int) (err error) {
	key := p.GetKeyDBUserMessage(data.UID)
	field := p.GetFieldDBUserMessage(data.SessionKey)
	setLRUCache(key, field, int64(ttl), data)
	return
}

func (p *Pika) SetEXDBUserMessageWithSess(ctx context.Context, data *pb.DBUserMessage, ttl int) (err error) {
	return p.SetEXDBUserMessage(ctx, data, ttl)
}

// DeleteDBUserMessage is generated from "//@db:hash|Pika|DBUserMessage:%s,UID|Session:%s,SessionKey"
//   err  error any error when put data to database
// if monitor is enabled on init this operation records prometheus
// metric: QPS and OperationProcessTime
func (p *Pika) DeleteDBUserMessage(ctx context.Context, UID string, SessionKey string) (ok bool, err error) {
	key := p.GetKeyDBUserMessage(UID)
	field := p.GetFieldDBUserMessage(SessionKey)
	delLRUCache(key, field)
	return
}
