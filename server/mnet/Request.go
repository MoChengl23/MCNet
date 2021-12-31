package mnet

import (
	"net"
	"server/face"
)

type Request struct {
	message []byte
	conn    net.Conn //有可能要用
	session face.ISession
}

func (request *Request) GetMessage() []byte {
	return request.message
}

//发送回给自己
func (request *Request) SendMessage(isInRoom bool) {
	if isInRoom {
		request.GetSession().GetRoom().Broadcast(request.message)
	} else {
		request.session.SendMessage(request.GetMessage())
	}

}

func (request *Request) GetSession() face.ISession {
	return request.session
}
func (request *Request) GetConn() net.Conn {
	return request.conn
}
