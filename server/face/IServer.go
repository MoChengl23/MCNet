package face

type IServer interface {
	Start()
	Stop()
	Serve()

	GenerateNewRoom(players [3]uint32)
	SendMessageToClient(sid uint32, data []byte)
}
