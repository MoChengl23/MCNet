package face

import (
	"server/pb"
)

type IMatchMessageHandle interface {
	ResponseMatch(sid uint32, message *pb.PbMessage)
}
