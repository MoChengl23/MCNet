package room

import (
	"fmt"
	"server/face"
	"server/pb"
	"time"
)

type RoomStateSelect struct {
	room      face.IRoom
	readyArr  []bool
	selectArr []int
}

func (state *RoomStateSelect) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.readyArr[i] = false
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
	case <-time.After(time.Second * 999):
		fmt.Println("room confirm reachtime ")
		state.Dismiss()
		return
	}

}

func (state *RoomStateSelect) Update(sid uint32, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(sid)

	state.readyArr[index] = mes.SelectData.IsReady
	state.selectArr[index] = int(mes.SelectData.Faction)
	// if !mes.SelectData.IsReady{

	// }

	// selectdata := pb.MakeRoomSelectData(mes)
	state.room.Broadcast(pb.Byte(mes))

	// mes := pb.MakeSelectData()
	if state.CheckAllSelect() {
		state.room.SetSelectData(state.selectArr)
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
