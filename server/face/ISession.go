package face

import "net"

// "net"

type ISession interface {
	Start()
	Stop()
	GetConnection() net.Conn
	GetSid() uint32
	SetRoomId(roomId uint32) error
	GetRoomId() uint32
	GetRemoteAddress() string
	SendMessage(data []byte)
}
