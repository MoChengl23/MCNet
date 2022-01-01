package room

import (
	"fmt"
	"server/face"
	"server/pb"

	// "server/mnet"
	"sync"

	"google.golang.org/protobuf/proto"
)

type RoomState int32

const (
	roomStateConfirm      RoomState = 0
	roomStateSelect       RoomState = 1
	roomStateLoadResource RoomState = 2
	roomStateFight        RoomState = 3
	roomStateEnd          RoomState = 4
	roomStateNone         RoomState = 5
)

type Room struct {
	server   face.IServer
	stateMap map[RoomState]face.IRoomState
	stateId  RoomState
	roomId   uint32
	// playerSessions map[face.ISession]bool
	players [3]uint32
	// confirmArr []int32
	// selectArr []int32
	lock sync.Mutex
}

func (room *Room) Init() {
	//Add StateMap
	room.stateMap[roomStateConfirm] = &RoomStateConfirm{room, []bool{}}
	room.stateMap[roomStateSelect] = &RoomStateSelect{room}
	room.stateMap[roomStateLoadResource] = &RoomStateLoadResource{room}
	room.stateMap[roomStateFight] = &RoomStateFight{room}
	room.stateMap[roomStateEnd] = &RoomStateEnd{room}

	//向玩家发送进入房间命令，和房间成员信息
	mes := &pb.PbMessage{
		Cmd: pb.PbMessage_roomInit,
		RoomPlayersData: &pb.RoomPlayersData{
			PlayerData: []*pb.PlayerData{
				{PlaySid: 1,
					RoomId:      3,
					PlayerName:  "mocheng",
					PlayerIndex: 0,
					IsConfirmed: false,
					PlayerSelectData: &pb.PlayerSelectData{
						Faction:       1,
						IsSelectDone:  false,
						AllSelectDone: false,
					},
					IsReady: false,
					// LoadPercent:        99,
					IsLoadResourceDone: false,
				},
			},
		},
	}
	if data, err := proto.Marshal(mes); err != nil {
		fmt.Println(err)
	} else {
		room.Broadcast(data)
	}

}

func (room *Room) Stop() {

}

// func (room *Room) AddPlayer(session face.ISession) {
// 	room.lock.Lock()
// 	room.playerSessions = append(room.playerSessions, session)
// 	room.lock.Unlock()

// }
// func (room *Room) RemovePlayer(session face.ISession) {
// 	room.lock.Lock()
// 	for index, i := range room.playerSessions {
// 		if i == session {
// 			room.playerSessions = append(room.playerSessions[:index], room.playerSessions[index+1:]...)
// 		}
// 	}
// 	room.lock.Unlock()

// }

// func (room *Room) Broadcast(message *pb.PbMessage) {

// 	for _, sid := range room.players {
// 		room.server.SendMessageToClient(sid, data)
// 	}

// }

func (room *Room) Broadcast(data []byte) {
	for _, sid := range room.players {
		room.server.SendMessageToClient(sid, data)
	}

}

func (room *Room) ChangeRoomState(newState RoomState) {
	if room.stateId != newState {
		room.stateMap[room.stateId].Exit()
		//把所有状态设为false 供下一轮状态判断
		// for i, _ := range room.playerSessions {
		// 	room.playerSessions[i] = false
		// }

	}
	room.stateId = newState
	room.stateMap[room.stateId].Enter()

}

func (room *Room) SendConfirm(session face.ISession) {
	if room.stateId == roomStateConfirm {
		// room.GetState().UpdateConfirm(session)

	}
}

func (room *Room) SendSelect() {

}

func (room *Room) SendLoadResource() {

}

func (room *Room) SendFightStart() {

}

func (room *Room) GetRoomId() uint32 {
	return room.roomId
}
func (room *Room) GetRoomPlayerCount() int {
	return len(room.players)
}

func (room *Room) GetState() face.IRoomState {
	return room.stateMap[room.stateId]
}

//估计要删
// func (room *Room) SetState(roomState face.IRoomState) {
// 	if room.roomState != roomState {
// 		room.roomState = roomState
// 	}

// }

func NewRoom(_server face.IServer, players [3]uint32) *Room {
	newRoom := &Room{
		server:   _server,
		stateMap: make(map[RoomState]face.IRoomState),
		stateId:  roomStateNone,
		roomId:   players[0],
		players:  players,
		lock:     *new(sync.Mutex),
	}
	newRoom.Init()
	return newRoom
}

// func AddPlayer(players []uint32) map[face.ISession]boo {
// 	sessionMap := make(map[face.ISession]bol)
// 	for _, session := range playerSession {
// 		sessionMap[session] = fase
//	}
// 	return sessionap
