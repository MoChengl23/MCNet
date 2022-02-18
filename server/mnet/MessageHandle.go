package mnet

import (
	"fmt"
	"server/face"
	match "server/mnet/Match"
	"server/mnet/room"
	"server/pb"

	"google.golang.org/protobuf/proto"
)

type WorkerPool struct {
	server         face.IServer
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest

	matchMessageHandle face.IMessageHandle
	roomMessageHandle  face.IMessageHandle
}

func (workerPool *WorkerPool) Init() {
	fmt.Println("WorkerPool Init")
	workerPool.StartWorkerPool()
	workerPool.matchMessageHandle = match.NewMatchMessageHandle(workerPool.server.GetMatchSystem())
	workerPool.roomMessageHandle = room.NewRoomMessageHandle(workerPool.server)

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
		workerPool.matchMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_room:
		workerPool.roomMessageHandle.Response(request.GetSession(), mes)
		// case pb.PbMessage_fight:
		// 	workerPool.ResponseTest(request.GetSid(), mes)
	}
}

func (workerPool *WorkerPool) ResponseLogin(sid uint32) {
	mes := pb.MakeLogin()

	workerPool.server.SendMessageToClient(sid, mes)

}
//test
func (workerPool *WorkerPool) ResponseTest(sid uint32, mes1 *pb.PbMessage) {
	mes := pb.Byte(mes1)

	for sid := range workerPool.server.GetAllPlayer() {
		workerPool.server.SendMessageToClient(sid, mes)
	}

}

func (workerPool *WorkerPool) StartWorkerPool() {
	for i := 0; i < int(workerPool.WorkerPoolSize); i++ {
		workerPool.TaskQueue[i] = make(chan face.IRequest)
		go workerPool.StartOneWorker(i, workerPool.TaskQueue[i])
	}

}
func (workerPool *WorkerPool) StartOneWorker(workerID int, taskQueue chan face.IRequest) {
	fmt.Println("WorkerId = ", workerID, "  Start")
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

func NewWorkerPool(_server face.IServer) *WorkerPool {
	return &WorkerPool{
		server:         _server,
		WorkerPoolSize: 10,
		TaskQueue:      make([]chan face.IRequest, 10),
		// roomMessageHandle: &room.RoomMessageHandle{},
	}

}
