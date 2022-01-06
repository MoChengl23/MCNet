package mnet

import (
	"server/face"
)

type Request struct {
	message []byte
	roomId  uint32
	sid     uint32
	session face.ISession
}

func (request *Request) GetMessage() []byte {
	return request.message
}

func (request *Request) GetSession() face.ISession {
	return request.session
}
func (request *Request) GetSid() uint32 {
	return request.sid
}
func (request *Request) GetRoomId() uint32 {

	return request.roomId
}
