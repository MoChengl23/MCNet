package roomSystem

import (
	"fmt"
	"server/face"
	"server/pb"
)

type RoomMessageHandle struct {
	server face.IServer
}
func (messageHandle *RoomMessageHandle) Init(server face.IServer){
	messageHandle.server = server
}

func (messageHandle *RoomMessageHandle) Response(session face.ISession, message *pb.PbMessage) {
	// sid := session.GetSid()
	roomId := session.GetCurrentRoomId()
	if roomId == 0 {
		return

	}
	fmt.Println("Receive :", message)

	room := messageHandle.server.GetRoom(roomId)
	if room != nil {

		room.GetCurrentState().Update(session, message)
	} else {
		fmt.Println("No Room")
	}

}

func NewRoomMessageHandle(server face.IServer) face.IMessageHandle {
	return &RoomMessageHandle{
		server: server,
	}

}
