package face

import "net"

//判断此时玩家在对局里还是在大厅里
type IRequest interface {
	GetMessage() []byte
	GetSession() ISession
	GetConn() net.Conn
	SendMessage(isInRoom bool)
}
