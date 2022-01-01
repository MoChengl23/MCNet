package face

import (
	"server/pb"
)

type IMatchMessageHandle interface {
	ResponseMatch(message *pb.PbMessage)
}
