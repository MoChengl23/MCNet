package face

import "server/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(sid uint32, mes *pb.PbMessage)
}
