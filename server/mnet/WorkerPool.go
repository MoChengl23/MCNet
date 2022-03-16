package mnet

import (
	"fmt"
	"server/face"
	"server/mnet/matchSystem"
	"server/mnet/roomSystem"
	"server/pb"
	"server/singleton"

	"google.golang.org/protobuf/proto"
)

type WorkerPool struct {
	server         face.IServer
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest

	matchMessageHandle face.IMessageHandle
	roomMessageHandle  face.IMessageHandle
}

func (workerPool *WorkerPool) Init(server face.IServer) {
	fmt.Println("WorkerPool Init")
	workerPool.server = server
	workerPool.WorkerPoolSize = 10

	workerPool.TaskQueue = make([]chan face.IRequest, 10)

	workerPool.StartWorkerPool()
	singleton.Singleton[matchSystem.MatchMessageHandle]().Init(workerPool.server)
	singleton.Singleton[roomSystem.RoomMessageHandle]().Init(workerPool.server)
	// workerPool.matchMessageHandle = singleton.Singleton[matchSystem.MatchMessageHandle]()
	// workerPool.roomMessageHandle = singleton.Singleton[roomSystem.RoomMessageHandle]()
	// workerPool.matchMessageHandle = matchSystem.NewMatchMessageHandle(workerPool.server.GetMatchSystem())
	// workerPool.roomMessageHandle = roomSystem.NewRoomMessageHandle(workerPool.server)

}

func (workerPool *WorkerPool) DoMessageHandler(request face.IRequest) {

	//测试下Pb能不能解码
	mes := &pb.PbMessage{}
	if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		fmt.Println(err)
	}

	switch mes.Cmd {
	case pb.PbMessage_login:
		workerPool.ResponseLogin(request.GetSession().GetSid())

	case pb.PbMessage_match:
		singleton.Singleton[matchSystem.MatchMessageHandle]().Response(request.GetSession(), mes)
		// workerPool.matchMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_room:
		singleton.Singleton[roomSystem.RoomMessageHandle]().Response(request.GetSession(), mes)
		// workerPool.roomMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_chat:
		workerPool.ResponseTest(request.GetSession())
	}
}

func (workerPool *WorkerPool) ResponseLogin(sid uint32) {
	mes := pb.MakeLogin()

	workerPool.server.SendMessageToClient(sid, mes)

}

//test
func (workerPool *WorkerPool) ResponseTest(session face.ISession) {
	// mes := pb.Byte(mes1)

	// for sid := range workerPool.server.GetAllPlayer() {
	// 	workerPool.server.SendMessageToClient(sid, mes)
	// }

}

func (workerPool *WorkerPool) StartWorkerPool() {

	for i := 0; i < int(workerPool.WorkerPoolSize); i++ {
		workerPool.TaskQueue[i] = make(chan face.IRequest)
		go workerPool.StartOneWorker(i, workerPool.TaskQueue[i])
	}

}
func (workerPool *WorkerPool) StartOneWorker(workerID int, taskQueue chan face.IRequest) {

	// fmt.Println("WorkerId = ", workerID, "  Start")
	for {
		request := <-taskQueue
		workerPool.DoMessageHandler(request)
		fmt.Println(workerID, "work over")

	}
}
func (workerPool *WorkerPool) AddToTaskQueue(request face.IRequest) {

	workerID := request.GetSession().GetSid() % workerPool.WorkerPoolSize
	fmt.Println("AddTaskQueue  ", workerID)

	workerPool.TaskQueue[workerID] <- request

}

// func NewWorkerPool(_server face.IServer) *WorkerPool {
// 	return &WorkerPool{
// 		server:         _server,
// 		WorkerPoolSize: 10,
// 		TaskQueue:      make([]chan face.IRequest, 10),
// 		// roomMessageHandle: &room.RoomMessageHandle{},
// 	}

// }
