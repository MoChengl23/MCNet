package pb

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func Byte(mes *PbMessage) []byte {
	if data, err := proto.Marshal(mes); err != nil {
		fmt.Println(err)
	} else {
		return data
	}
	return nil
}

func MakeLoginMessage() []byte {

	mes := &PbMessage{
		Cmd:  PbMessage_login,
		Name: "login",
	}
	return Byte(mes)

}
func MakeJoinMatch() []byte {
	mes := &PbMessage{
		Cmd:      PbMessage_match,
		CmdMatch: PbMessage_joinMatch,
	}
	return Byte(mes)
}
func MakeQuitMatch() []byte {
	mes := &PbMessage{
		Cmd:      PbMessage_match,
		CmdMatch: PbMessage_quitMatch,
	}
	return Byte(mes)
}

//server 发给client: Cmd,CmdRoom: PbMessage_confirm

//client 发给server: Cmd,CmdRoom: PbMessage_confirm, sid

func MakeRoomConfirmMessage() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_confirm,
	}
	return Byte(mes)
}
func MakeRoomIndex(index int32) []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_confirm,
		Index: index,
	}
	return Byte(mes)
}

func MakeRoomDismissMes() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_dismissed,
	}
	return Byte(mes)

}

// func MakeSelectData( int) []byte{
// 	mes := &PbMessage{

// 	}
// }

func MakeRoomSelectMessage() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_select,
	}
	return Byte(mes)
}

func MakeRoomSelectData(mes *PbMessage){

}
