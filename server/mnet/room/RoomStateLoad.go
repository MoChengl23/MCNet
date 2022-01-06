package room

import (
	"fmt"
	"server/face"
	"server/pb"
)

type RoomStateLoadResource struct {
	room face.IRoom
}

func (state *RoomStateLoadResource) Enter() {
	fmt.Println("进入load阶段")

}
func (state *RoomStateLoadResource) Exit() {

}

func (state *RoomStateLoadResource) Update(sid uint32, mes *pb.PbMessage) {

}
