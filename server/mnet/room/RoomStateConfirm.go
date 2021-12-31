package room

// import "server/face"

type RoomStateConfirm struct {
	room       *Room
	confirmArr []bool
}

func (state *RoomStateConfirm) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.confirmArr[i] = false
	}

}
func (state *RoomStateConfirm) Exit() {

}

func (state *RoomStateConfirm) Update() {

}

func (state *RoomStateConfirm) UpdateConfirm(index int) {
	state.confirmArr[index] = true
	if state.CheckAllConfirm() {
		// state.room.ChangeRoomState()
	}

}
func (state *RoomStateConfirm) CheckAllConfirm() bool {
	for _, i := range state.confirmArr {
		if i == false {
			return false
		}
	}
	return true

}
