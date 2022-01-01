package face

type IServer interface {
	Start()
	Stop()
	Serve()

	GenerateNewRoom(players [1]uint32)
	GetMatchSystem() IMatchSystem
	SendMessageToClient(sid uint32, data []byte)
}
