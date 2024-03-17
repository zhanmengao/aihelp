package cache

// DB: Pika, TotalKeys: 2
const (
	keyDBChatSession   = "DBUser:%s" // type: hash
	fieldDBChatSession = "DBChatSession:%s"
	keyDBUserMessage   = "DBUser:%s" // type: hash
	fieldDBUserMessage = "DBUserMessage"
)

type Pika struct{}
