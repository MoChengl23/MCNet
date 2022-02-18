package face

import (
	"server/pb"
)

type IMessageHandle interface {
	Response(session ISession, message *pb.PbMessage)
}
