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

func (request *Request) GetSession() face.ISession {
	return request.session
}
func (request *Request) GetConn() net.Conn {
	return request.conn
}
