package room

import "server/face"

type RoomStateEnd struct {
	room face.IRoom
}

func (state *RoomStateEnd) Enter() {

}
func (state *RoomStateEnd) Exit() {

}

func (state *RoomStateEnd) Update() {

}
