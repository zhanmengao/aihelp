//
// +//
package pb

import (
	"fmt"
	redisdb"github.com/zhanmengao/aihelp/database"
)

func NewDBUserMessage(UID string,SessionKey string,)*DBUserMessage{
    ret := &DBUserMessage{
    UID : UID,
    SessionKey : SessionKey,
    }
    return ret
}

func (p *DBUserMessage) GetDBKey() string {
	key := fmt.Sprintf("DBUserMessage:%s", p.UID, )
	return key
}

func(p *DBUserMessage) GetDBKeyFormat() string {
    return "DBUserMessage:%s"
}

func (p *DBUserMessage) GetDBField() string {
	key := fmt.Sprintf("Session:%s", p.SessionKey, )
	return key
}

func(p *DBUserMessage) GetDBFieldFormat() string {
    return "Session:%s"
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
