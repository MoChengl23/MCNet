package face

import (
	"server/pb"
)

type IMatchMessageHandle interface {
	ResponseMatch(roomId uint32, message *pb.PbMessage)


}
