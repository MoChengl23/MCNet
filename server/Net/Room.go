package mnet

import "server/face"

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
	roomState     RoomState
	roomId        uint32
	playerSessions []face.ISession
}

func (room *Room) Start() {

}

func (room *Room) Stop() {

}

func (room *Room) AddPlayer(session face.ISession) {
	// server.sessionMap[playerSid].SetRoomId(roomId)
	room.playerSessions = append(room.playerSessions, session)

}

func (room *Room) Broadcast(data []byte) {
	for _,player := range room.playerSessions {
			player.SendMessage(data)
	}

}

func (room *Room) GetRoomId() uint32 {
	return room.roomId
}

func (room *Room) GetState() RoomState {
	return room.roomState
}
func (room *Room) SetState(roomState RoomState) {
	room.roomState = roomState
}
