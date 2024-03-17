//
// +//
package pb

import (
	"fmt"
	redisdb"github.com/zhanmengao/aihelp/database"
)

func NewDBChatSession(UID string,SessionKey string,)*DBChatSession{
    ret := &DBChatSession{
    UID : UID,
    SessionKey : SessionKey,
    }
    return ret
}

func (p *DBChatSession) GetDBKey() string {
	key := fmt.Sprintf("DBUser:%s", p.UID, )
	return key
}

func(p *DBChatSession) GetDBKeyFormat() string {
    return "DBUser:%s"
}

func (p *DBChatSession) GetDBField() string {
	key := fmt.Sprintf("DBChatSession:%s", p.SessionKey, )
	return key
}

func(p *DBChatSession) GetDBFieldFormat() string {
    return "DBChatSession:%s"
}

func(p *DBChatSession)GetDBKeyName()string{
    return "DBChatSession"
}



func(p *DBChatSession)GetDBName()string{
    return "Pika"
}

func(p *DBChatSession)CopyFrom(other redisdb.IDBProto)(err error){
    body, err := other.Marshal()
    if err!=nil{
        return
    }
    err = p.Unmarshal(body)
    if err!=nil{
        return
    }
    return
}

func NewDBUserMessage(UID string,)*DBUserMessage{
    ret := &DBUserMessage{
    UID : UID,
    }
    return ret
}

func (p *DBUserMessage) GetDBKey() string {
	key := fmt.Sprintf("DBUser:%s", p.UID, )
	return key
}

func(p *DBUserMessage) GetDBKeyFormat() string {
    return "DBUser:%s"
}

func (p *DBUserMessage) GetDBField() string {
	key := fmt.Sprintf("DBUserMessage", )
	return key
}

func(p *DBUserMessage) GetDBFieldFormat() string {
    return "DBUserMessage"
}

func(p *DBUserMessage)GetDBKeyName()string{
    return "DBUserMessage"
}



func(p *DBUserMessage)GetDBName()string{
    return "Pika"
}

func(p *DBUserMessage)CopyFrom(other redisdb.IDBProto)(err error){
    body, err := other.Marshal()
    if err!=nil{
        return
    }
    err = p.Unmarshal(body)
    if err!=nil{
        return
    }
    return
}
