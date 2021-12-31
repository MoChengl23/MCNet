package face

import "server/pb"

type IMatchSystem interface {
	Init()
	EnterMatchQueue(sid uint32)
	QuitMatchQueue(sid uint32)
	ResponseMatch(message *pb.PbMessage)
}
