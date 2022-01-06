package face

type IServer interface {
	Start()
	Stop()
	Serve()

	GenerateNewRoom(players [3]uint32)
	GetRoom(roomId uint32) IRoom
	AddRoom(roomId uint32, room IRoom)


	GetPlayerSession(sid uint32) ISession
	DeleteRoom(roomId uint32)
	GetMatchSystem() IMatchSystem
	SendMessageToClient(sid uint32, data []byte)
}
