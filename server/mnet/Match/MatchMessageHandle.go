package match

import (
	"fmt"
	"server/face"
	"server/pb"
)

type MatchMessageHandle struct {
	matchSystem face.IMatchSystem
}

func (matchMessageHandle *MatchMessageHandle) ResponseMatch(message *pb.PbMessage) {
	fmt.Println("a player join match", message.Sid)
	matchMessageHandle.matchSystem.UpdateMatchQueue(message)
}

func NewMatchMessageHandle(matchSystem face.IMatchSystem) *MatchMessageHandle {
	return &MatchMessageHandle{
		matchSystem: matchSystem,
	}
}
