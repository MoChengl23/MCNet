package match

import (
	"container/list"
	"fmt"
	"server/face"
	"server/pb"
	"sync"
)

type MatchSystem struct {
	server     face.IServer
	matchQueue list.List

	lock *sync.Mutex
}

func (match *MatchSystem) Init() {
	fmt.Println("match System Init")
	fmt.Println("初始匹配队列长度  ", match.matchQueue.Len())
}

func (match *MatchSystem) UpdateMatchQueue(message *pb.PbMessage) {

	switch message.CmdMatch {
	case pb.PbMessage_joinMatch:
		match.EnterMatchQueue(message.Sid)

	case pb.PbMessage_quitMatch:
		match.QuitMatchQueue(message.Sid)

	default:

	}

}

func (match *MatchSystem) EnterMatchQueue(sid uint32) {
	match.lock.Lock()
	match.matchQueue.PushBack(sid)
	for i := match.matchQueue.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	match.lock.Unlock()

	if match.matchQueue.Len() >= 1 {
		match.GenerateNewRoom()
	}

}

func (match *MatchSystem) QuitMatchQueue(sid uint32) {
	match.lock.Lock()
	for i := match.matchQueue.Front(); i != nil; i = i.Next() {
		if i.Value == sid {
			match.matchQueue.Remove(i)
		}
	}
	match.lock.Unlock()
}

func (match *MatchSystem) GenerateNewRoom() {
	match.lock.Lock()
	roomPlayers := new([1]uint32)
	for i := 0; i < 1; i++ {
		//这里的断言可能有问题
		fmt.Println(match.matchQueue.Front().Value)

		roomPlayers[i] = match.matchQueue.Front().Value.(uint32)
		match.matchQueue.Remove(match.matchQueue.Front())
	}
	match.server.GenerateNewRoom(*roomPlayers)
	match.lock.Unlock()

}

func (match *MatchSystem) Update() {

}
func NewMatchSystem(_server face.IServer) *MatchSystem {
	return &MatchSystem{
		server: _server,
		// matchQueue: *list.New(),
		lock: new(sync.Mutex),
	}

}
