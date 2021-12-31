package room

import (
	"server/face"
	"server/pb"
)

type RoomMessageHandle struct {
	messageHandle face.IMessageHandle
}

func (messageHandle *RoomMessageHandle) ResponseLogin(roomId uint32, message *pb.PbMessage) {

}

func (messageHandle *RoomMessageHandle) ResponseMatch(roomId uint32, message *pb.PbMessage) {

}
func (messageHandle *RoomMessageHandle) ResponseConfirm(roomId uint32, message *pb.PbMessage) {

}
func (messageHandle *RoomMessageHandle) ResponseRoomInit(roomId uint32, message *pb.PbMessage) {

}
func (messageHandle *RoomMessageHandle) ResponseSelect(roomId uint32, message *pb.PbMessage) {

}
func (messageHandle *RoomMessageHandle) ResponseLoadResource(roomId uint32, message *pb.PbMessage) {

}
func (messageHandle *RoomMessageHandle) ResponseFightStart(roomId uint32, message *pb.PbMessage) {

}

func NewRoomMessageHandle(messageHandle face.IMessageHandle) face.IRoomMessageHandle {
	return &RoomMessageHandle{
		messageHandle: messageHandle,
	}

}
