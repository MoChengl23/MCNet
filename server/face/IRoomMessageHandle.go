package face

import (
	"server/pb"
)

type IRoomMessageHandle interface {
	ResponseRoom(sid uint32,roomId uint32, message *pb.PbMessage)

	// ResponseConfirm(roomId uint32, message *pb.PbMessage)
	// ResponseRoomInit(roomId uint32, message *pb.PbMessage)
	// ResponseSelect(roomId uint32, message *pb.PbMessage)
	// ResponseLoadResource(roomId uint32, message *pb.PbMessage)
	// ResponseFightStart(roomId uint32, message *pb.PbMessage)
}
