package face

// import "mnet"

type IRoom interface {
	Init()
	Stop()
	Broadcast(data []byte)
	GetState() IRoomState
	// SetState(roomState  IRoomState)

}
