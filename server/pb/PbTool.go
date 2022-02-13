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

func MakeLogin() []byte {

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

func MakeRoomConfirm() []byte {
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
		Index:   index,
	}
	return Byte(mes)
}

func MakeRoomDismiss() []byte {
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

func MakeRoomSelect() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_select,
	}
	return Byte(mes)
}

func MakeRoomSelectData(mes *PbMessage) {

}

func MakeRoomLoadCmd() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_load,
	}
	return Byte(mes)
}
func MakeRoomLoadData(loadPercent int32) []byte {
	mes := &PbMessage{
		Cmd:         PbMessage_room,
		CmdRoom:     PbMessage_loadData,
		LoadPercent: loadPercent,
	}
	return Byte(mes)
}

func MakeFightStartCmd() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_fightStart,
	}
	return Byte(mes)
}

func MakeFightData(frameId int, fightMessage []*FightMessage) []byte {
	mes := &PbMessage{
		Cmd:          PbMessage_room,
		CmdRoom:      PbMessage_fightOp,
		FrameId:      int32(frameId),
		FightMessage: fightMessage,
	}
	return Byte(mes)
}
func MakeFightEnd() []byte {
	mes := &PbMessage{
		Cmd: PbMessage_fightEnd,
	}
	return Byte(mes)
}
