package face

import (
	"server/pb"
)

type IMessageHandle interface {
	Init(server IServer)
	Response(session ISession, message *pb.PbMessage)
}
