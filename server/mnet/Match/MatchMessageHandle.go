package match

import (
	"server/face"
	"server/pb"
)

type MatchMessageHandle struct {
	matchSystem face.IMatchSystem
}

func (matchMessageHandle *MatchMessageHandle) Response(session face.ISession, message *pb.PbMessage) {
	sid := session.GetSid()
	matchMessageHandle.matchSystem.UpdateMatchQueue(message, sid)
}

func NewMatchMessageHandle(matchSystem face.IMatchSystem) face.IMessageHandle {
	return &MatchMessageHandle{
		matchSystem: matchSystem,
	}
}
