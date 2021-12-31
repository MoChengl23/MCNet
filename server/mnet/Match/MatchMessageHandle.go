package match

import (
	"server/face"
	"server/pb"
)

type MatchMessageHandle struct {
	server face.IServer
}

func (matchMessageHandle *MatchMessageHandle) ResponseMatch(roomId uint32, message *pb.PbMessage) {

}
