package mnet

import (
	"fmt"
	"server/face"
	match "server/mnet/Match"
	"server/mnet/room"
	"server/pb"

	"google.golang.org/protobuf/proto"
)

type MessageHandle struct {
	server         face.IServer
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest

	matchMessageHandle face.IMatchMessageHandle
	roomMessageHandle  face.IRoomMessageHandle
}

func (messageHandle *MessageHandle) Init() {
	fmt.Println("MessageHandle Init")
	messageHandle.StartWorkerPool()
	messageHandle.matchMessageHandle = match.NewMatchMessageHandle(messageHandle.server.GetMatchSystem())
	messageHandle.roomMessageHandle = room.NewRoomMessageHandle(messageHandle.server)

}

func (messageHandle *MessageHandle) DoMessageHandler(request face.IRequest) {
	fmt.Println("DoMesagehandle")

	//测试下Pb能不能解码
	mes := &pb.PbMessage{}
	if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		fmt.Println(err)
	}
	fmt.Println("收到的信息是 ", mes.Cmd)

	switch mes.Cmd {
	case pb.PbMessage_login:
		messageHandle.ResponseLogin(request.GetSid())

	case pb.PbMessage_match:
		messageHandle.matchMessageHandle.ResponseMatch(request.GetSid(), mes)
	case pb.PbMessage_room:
		messageHandle.roomMessageHandle.ResponseRoom(request.GetSid(), request.GetRoomId(), mes)

	}
}

func (messageHandle *MessageHandle) ResponseLogin(sid uint32) {
	mes := pb.MakeLoginMessage()
	fmt.Println(" response login")
	messageHandle.server.SendMessageToClient(sid, mes)

}


func (messageHandle *MessageHandle) StartWorkerPool() {
	for i := 0; i < int(messageHandle.WorkerPoolSize); i++ {
		messageHandle.TaskQueue[i] = make(chan face.IRequest)
		go messageHandle.StartOneWorker(i, messageHandle.TaskQueue[i])
	}

}
func (messageHandle *MessageHandle) StartOneWorker(workerID int, taskQueue chan face.IRequest) {
	fmt.Println("WorkerId = ", workerID, "  Start")
	for {
		select {
		case request := <-taskQueue:

			messageHandle.DoMessageHandler(request)
		}
	}
}
func (messageHandle *MessageHandle) AddToTaskQueue(request face.IRequest) {

	workerID := request.GetSession().GetSid() % messageHandle.WorkerPoolSize
	fmt.Println("AddTaskQueue  ", workerID)
	// workerID = uint32(rand.Intn(10))
	// fmt.Println("AddTaskQueue  ", workerID)
	// messageHandle.TaskQueue[workerID] <- request
	//消息队列有bug，先暂时这么处理
	go messageHandle.DoMessageHandler(request)

}

func NewMessageHandler(_server face.IServer) *MessageHandle {
	return &MessageHandle{
		server:         _server,
		WorkerPoolSize: 10,
		TaskQueue:      make([]chan face.IRequest, 10),
		// roomMessageHandle: &room.RoomMessageHandle{},
	}

}
