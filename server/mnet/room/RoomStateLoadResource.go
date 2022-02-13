package room

import (
	"fmt"
	"server/face"
	"server/pb"
)

type RoomStateLoadResource struct {
	room     face.IRoom
	percent  []int
	loadDone []bool
}

func NewRoomStateLoadResource(room face.IRoom, length int) face.IRoomState {
	state := &RoomStateLoadResource{
		room:     room,
		percent:  make([]int, length),
		loadDone: make([]bool, length),
	}
	return state
}

func (state *RoomStateLoadResource) Enter() {
	fmt.Println("进入load阶段")

	mes := pb.MakeRoomLoadCmd()
	state.room.Broadcast(mes)

}
func (state *RoomStateLoadResource) Exit() {

}

func (state *RoomStateLoadResource) Update(sid uint32, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(sid)
	loadPercent := mes.LoadPercent
	if loadPercent == 100 {
		state.loadDone[index] = true
		if state.CheckAllLoadDone() {
			state.room.ChangeRoomState(int(roomStateFight))
		}
	} else {
		state.percent[index] = int(loadPercent)
		state.room.Broadcast(pb.Byte(mes))
	}

}
func (state *RoomStateLoadResource) CheckAllLoadDone() bool {
	return true
}
