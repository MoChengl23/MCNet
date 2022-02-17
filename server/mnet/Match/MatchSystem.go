package match

import (
	"container/list"
	"fmt"
	"server/face"
	"server/mnet/room"
	"server/pb"
	"sync"
)

type MatchSystem struct {
	server     face.IServer
	matchQueue list.List
	matchlen   int
	lock       *sync.Mutex
}

func (match *MatchSystem) Init() {
	// fmt.Println("match System Init")
	// fmt.Println("初始匹配队列长度  ", match.matchQueue.Len())
}

func (match *MatchSystem) UpdateMatchQueue(message *pb.PbMessage, sid uint32) {

	switch message.CmdMatch {
	case pb.PbMessage_joinMatch:
		match.EnterMatchQueue(sid)

	case pb.PbMessage_quitMatch:
		match.QuitMatchQueue(sid)

	}

}

func (match *MatchSystem) EnterMatchQueue(sid uint32) {
	fmt.Println("a player join match", sid)
	match.lock.Lock()
	match.matchQueue.PushBack(sid)

	match.lock.Unlock()
	
	mes := pb.MakeJoinMatch()
	match.server.SendMessageToClient(sid, mes)

	if match.matchQueue.Len() >= match.matchlen {
		match.GenerateNewRoom()
	}

}

func (match *MatchSystem) QuitMatchQueue(sid uint32) {
	fmt.Println("a player quit match", sid)
	match.lock.Lock()
	for i := match.matchQueue.Front(); i != nil; i = i.Next() {
		if i.Value == sid {
			match.matchQueue.Remove(i)
		}
	}
	match.lock.Unlock()
	mes := pb.MakeQuitMatch()
	match.server.SendMessageToClient(sid, mes)
}

func (match *MatchSystem) GenerateNewRoom() {

	match.lock.Lock()

	roomPlayers := make([]uint32, match.matchlen)
	for i := 0; i < match.matchlen; i++ {
		sid := match.matchQueue.Front().Value.(uint32)
		roomPlayers[i] = sid
		match.matchQueue.Remove(match.matchQueue.Front())
	}

	newRoom := room.NewRoom(match.server, roomPlayers)

	newRoom.Init()
	fmt.Println("newww rom")

	match.lock.Unlock()

}

func (match *MatchSystem) Update() {

}
func NewMatchSystem(_server face.IServer) *MatchSystem {
	return &MatchSystem{
		server:   _server,
		matchlen: 1,
		lock:     new(sync.Mutex),
	}

}
