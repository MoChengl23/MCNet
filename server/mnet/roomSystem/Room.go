package roomSystem

import (
	"server/face"

	// "server/pb"

	// "server/mnet"

	"sync"
)

type RoomState int

const (
	roomStateConfirm      RoomState = 0
	roomStateSelect       RoomState = 1
	roomStateLoadResource RoomState = 2
	roomStateFight        RoomState = 3
	roomStateEnd          RoomState = 4
	roomStateNone         RoomState = 5
)

type SelectData struct {
	selectId int
	isReady  bool
}

//a single room
type Room struct {
	server       face.IServer
	stateMap     map[RoomState]face.IRoomState
	currentState RoomState
	roomId       uint32
	// playerSessions map[face.ISession]bool
	playerSessions []face.ISession
	selectArr      []int
	// playSelectData []SelectData
	// selectArr []int32
	lock sync.Mutex
}

func (room *Room) Init() {
	// room.server
	room.server.AddRoom(room.roomId, room)
	// room.server

	// fmt.Println("New Room Init")
	length := len(room.playerSessions)
	//Add StateMap
	room.stateMap[roomStateConfirm] = NewRoomStateConfirm(room, length)
	room.stateMap[roomStateSelect] = NewRoomStateSelect(room, length)
	room.stateMap[roomStateLoadResource] = NewRoomStateLoadResource(room, length)
	room.stateMap[roomStateFight] = NewRoomStateFight(room, length)
	// room.stateMap[roomStateEnd] = &RoomStateEnd{room}

	room.ChangeRoomState(int(roomStateConfirm))

}

func (room *Room) Delete() {
	room.server.RemoveRoom(room.roomId)

}

func (room *Room) Broadcast(data []byte) {
	if data != nil {
		for _, session := range room.playerSessions {
			session.SendMessage(data)
			// room.server.SendMessageToClient(sid, data)
		}
	}
}

// func (room *Room) SendIndex() {

// 	for index, sid := range room.playerSessions {
// 		mes := pb.MakeRoomIndex(int32(index))
// 		room.server.SendMessageToClient(sid, mes)
// 	}

// }

func (room *Room) ChangeRoomState(newState int) {

	if int(room.currentState) != newState {
		room.stateMap[room.currentState].Exit()

	}

	room.currentState = RoomState(newState)

	room.stateMap[room.currentState].Enter()

}
func (room *Room) ChangePlayersRoomId() {
	for _, session := range room.playerSessions {
		session.ChangeRoomId(room.roomId)
		// room.server.GetSession(i).ChangeRoomId(room.roomId)
	}
}

// func (room *Room) UpdateConfirm(mes *pb.PbMessage) {
// 	if room.currentState == roomStateConfirm {
// 		playerIndex := room.GetPlayerIndex(mes.SId)
// 		if playerIndex != -1 {
// 			room.stateMap[roomStateConfirm].Update(playerIndex)
// 		}

// 	}
// }

// func (room *Room) UpdateSelect(mes *pb.PbMessage) {
// 	if room.currentState == roomStateSelect {
// 		room.stateMap[roomStateSelect].Update(int(mes.PlayerData.PlayerIndex))
// 	}

// }

func (room *Room) SendLoadResource() {

}

func (room *Room) SendFightStart() {

}

func (room *Room) GetRoomId() uint32 {
	return room.roomId
}
func (room *Room) GetRoomPlayerCount() int {
	return len(room.playerSessions)
}
func (room *Room) GetPlayerIndex(session face.ISession) int {
	for index, i := range room.playerSessions {
		if i == session {
			return index
		}
	}
	return -1
}

func (room *Room) SetSelectData(selectArr []int) {
	room.selectArr = selectArr
}
func (room *Room) GetSelectData() []int {
	return room.selectArr
}
func (room *Room) GetCurrentState() face.IRoomState {
	return room.stateMap[room.currentState]
}

func NewRoom(_server face.IServer, playerSessions []face.ISession) *Room {

	newRoom := &Room{
		server:         _server,
		stateMap:       make(map[RoomState]face.IRoomState),
		currentState:   roomStateConfirm,
		roomId:         playerSessions[0].GetSid(),
		playerSessions: playerSessions,
		lock:           *new(sync.Mutex),
	}
	// newRoom.Init()
	return newRoom
}

// func AddPlayer(players []uint32) map[face.ISession]boo {
// 	sessionMap := make(map[face.ISession]bol)
// 	for _, session := range playerSession {
// 		sessionMap[session] = fase
//	}
// 	return sessionap
