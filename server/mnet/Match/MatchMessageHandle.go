package match

import (
	"server/face"
	"server/pb"
)

type MatchMessageHandle struct {
	matchSystem face.IMatchSystem
}

func (matchMessageHandle *MatchMessageHandle) ResponseMatch(sid uint32, message *pb.PbMessage) {

	matchMessageHandle.matchSystem.UpdateMatchQueue(message, sid)
}

func NewMatchMessageHandle(matchSystem face.IMatchSystem) *MatchMessageHandle {
	return &MatchMessageHandle{
		matchSystem: matchSystem,
	}
}
