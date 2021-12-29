package face

// import "mnet"

type IRoom interface {
	Start()
	Stop()
	Broadcast(data []byte)
	GetState() int32
	SetState(roomState int32)

	JoinRoom(session ISession)
	LeaveRoom(session ISession)
}
