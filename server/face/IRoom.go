package face

// type RoomState int32

// const (
// 	roomStateConfirm      RoomState = 0
// 	roomStateSelect       RoomState = 1
// 	roomStateLoadResource RoomState = 2
// 	roomStateFight        RoomState = 3
// 	roomStateEnd          RoomState = 4
// 	roomStateNone         RoomState = 5
// )

type IRoom interface {
	Init()
	Delete()
	Broadcast(data []byte)
	// BroadcastByte(data []byte)
	ChangeRoomState(newState int)
	ChangePlayersRoomId()
	GetState() IRoomState
	GetStateId() int
	GetPlayerIndex(sid uint32) int

	GetRoomPlayerCount() int
	GetRoomId() uint32

	// UpdateConfirm(mes *pb.PbMessage)
}
