package room

import (
	"fmt"
	"server/face"
	"server/pb"
	"time"
)

type RoomStateSelect struct {
	room     face.IRoom
	readyArr []bool
}

func (state *RoomStateSelect) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.readyArr = append(state.readyArr, false)
	}
	mes := pb.MakeRoomSelectMessage()
	state.room.Broadcast(mes)
	fmt.Println("RoomSelect")
	state.CheckTimeLimit()
}

func (state *RoomStateSelect) Dismiss() {

	mes := pb.MakeRoomDismissMes()
	state.room.Broadcast(mes)
	state.room.Delete()

}
func (state *RoomStateSelect) Exit() {

}

func (state *RoomStateSelect) CheckTimeLimit() {

	select {
	case <-time.After(time.Second * 60):
		fmt.Println("room confirm reachtime ")
		state.Dismiss()
		return
	}

}

func (state *RoomStateSelect) Update(sid uint32, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(sid)
	state.readyArr[index] = true

	// mes := pb.MakeSelectData()
	if state.CheckAllSelect() {
		state.room.ChangeRoomState(int(roomStateLoadResource))
	}
}

func (state *RoomStateSelect) CheckAllSelect() bool {
	for _, i := range state.readyArr {
		if !i {
			return false
		}
	}
	return true

}
