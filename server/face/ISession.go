package face

import "net"

// "net"

type ISession interface {
	Start()
	Stop()
	GetConnection() net.Conn
	GetSid() uint32
	IsInRoom() bool
	SetRoom(room IRoom)
	GetRoom() IRoom
	GetRemoteAddress() string
	SendMessage(data []byte)
}
