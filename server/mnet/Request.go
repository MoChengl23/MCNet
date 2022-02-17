package mnet

import (
	"server/face"
)

type Request struct {
	message []byte

	session face.ISession
}

func (request *Request) GetMessage() []byte {
	return request.message
}

func (request *Request) GetSession() face.ISession {
	return request.session
}
func (request *Request) GetSid() uint32 {
	return request.session.GetSid()
}
func (request *Request) GetRoomId() uint32 {

	return request.session.GetCurrentRoomId()
}

func NewRequest(message []byte, session face.ISession) face.IRequest {
	return &Request{
		message: message,
		session: session,
	}

}
