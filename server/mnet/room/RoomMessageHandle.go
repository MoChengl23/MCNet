package room

import (
	"fmt"
	"server/face"
	"server/pb"
)

type RoomMessageHandle struct {
	server face.IServer
}

func (messageHandle *RoomMessageHandle) ResponseRoom(sid uint32, roomId uint32, message *pb.PbMessage) {
	fmt.Println("当前房间号 ", roomId)
	if roomId == 0 {
		return

	}

	room := messageHandle.server.GetRoom(roomId)
	if room != nil {
		fmt.Println("这个房间的state是 ", room.GetStateId())
		room.GetState().Update(sid, message)
	} else {
		fmt.Println("No Room")
	}

}

// func (messageHandle *RoomMessageHandle) ResponseMatch(roomId uint32, message *pb.PbMessage) {

// }
// func (messageHandle *RoomMessageHandle) ResponseConfirm(roomId uint32, message *pb.PbMessage) {

// }
// func (messageHandle *RoomMessageHandle) ResponseRoomInit(roomId uint32, message *pb.PbMessage) {

// }
// func (messageHandle *RoomMessageHandle) ResponseSelect(roomId uint32, message *pb.PbMessage) {

// }
// func (messageHandle *RoomMessageHandle) ResponseLoadResource(roomId uint32, message *pb.PbMessage) {

// }
// func (messageHandle *RoomMessageHandle) ResponseFightStart(roomId uint32, message *pb.PbMessage) {

// }

func NewRoomMessageHandle(server face.IServer) face.IRoomMessageHandle {
	return &RoomMessageHandle{
		server: server,
	}

}
