package room

import (
	"fmt"
	"server/face"
	"server/pb"

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
type Room struct {
	server   face.IServer
	stateMap map[RoomState]face.IRoomState
	stateId  RoomState
	roomId   uint32
	// playerSessions map[face.ISession]bool
	playersid []uint32
	selectArr []int
	// playSelectData []SelectData
	// selectArr []int32
	lock sync.Mutex
}

func (room *Room) Init() {
	room.server.AddRoom(room.roomId, room)

	fmt.Println("New Room Init")
	length := len(room.playersid)
	//Add StateMap
	room.stateMap[roomStateConfirm] = &RoomStateConfirm{room, make([]bool, length)}
	room.stateMap[roomStateSelect] = &RoomStateSelect{room, make([]bool, length), make([]int, length)}
	room.stateMap[roomStateLoadResource] = &RoomStateLoadResource{room}
	// room.stateMap[roomStateFight] = &RoomStateFight{room}
	// room.stateMap[roomStateEnd] = &RoomStateEnd{room}

	room.ChangeRoomState(int(roomStateConfirm))

}

func (room *Room) Delete() {
	room.server.DeleteRoom(room.roomId)

}

func (room *Room) Broadcast(data []byte) {
	if data != nil {
		for _, sid := range room.playersid {
			room.server.SendMessageToClient(sid, data)
		}
	}
}
func (room *Room) SendIndex() {

	for index, sid := range room.playersid {
		mes := pb.MakeRoomIndex(int32(index))
		room.server.SendMessageToClient(sid, mes)
	}

}

func (room *Room) ChangeRoomState(newState int) {
	if int(room.stateId) != newState {
		room.stateMap[room.stateId].Exit()

	}
	room.stateId = RoomState(newState)
	room.stateMap[room.stateId].Enter()

}
func (room *Room) ChangePlayersRoomId() {
	for _, i := range room.playersid {
		room.server.GetPlayerSession(i).ChangeRoomId(room.roomId)
	}
}

// func (room *Room) UpdateConfirm(mes *pb.PbMessage) {
// 	if room.stateId == roomStateConfirm {
// 		playerIndex := room.GetPlayerIndex(mes.SId)
// 		if playerIndex != -1 {
// 			room.stateMap[roomStateConfirm].Update(playerIndex)
// 		}

// 	}
// }

// func (room *Room) UpdateSelect(mes *pb.PbMessage) {
// 	if room.stateId == roomStateSelect {
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
	return len(room.playersid)
}
func (room *Room) GetPlayerIndex(sid uint32) int {
	for index, i := range room.playersid {
		if i == sid {
			return index
		}
	}
	return -1
}

func (room *Room) SetSelectData(selectArr []int) {
	room.selectArr = selectArr
}

func (room *Room) GetState() face.IRoomState {
	return room.stateMap[room.stateId]
}
func (room *Room) GetStateId() int {
	return int(room.stateId)
}

//估计要删
// func (room *Room) SetState(roomState face.IRoomState) {
// 	if room.roomState != roomState {
// 		room.roomState = roomState
// 	}

// }

func NewRoom(_server face.IServer, playersid []uint32) *Room {
	newRoom := &Room{
		server:    _server,
		stateMap:  make(map[RoomState]face.IRoomState),
		stateId:   roomStateConfirm,
		roomId:    playersid[0],
		playersid: playersid,
		lock:      *new(sync.Mutex),
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
