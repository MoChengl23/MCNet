package mnet

import (
	"fmt"
	"server/face"
	"server/mnet/room"
	"server/pb"

	"google.golang.org/protobuf/proto"
)

type MessageHandle struct {
	server         face.IServer
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest

	matchSystem       face.IMatchSystem
	roomMessageHandle face.IRoomMessageHandle
}

func (messageHandle *MessageHandle) Init() {
	messageHandle.roomMessageHandle = room.NewRoomMessageHandle(messageHandle)
}

func (messageHandle *MessageHandle) DoMessageHandler(request face.IRequest) {

	//测试下Pb能不能解码
	mes := &pb.PbMessage{}
	if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mes)

	roomId := uint32(mes.PlayerData.RoomId)
	switch mes.Cmd {
	case pb.PbMessage_match:
		messageHandle.matchSystem.ResponseMatch(mes)
	case pb.PbMessage_confirm:
		messageHandle.roomMessageHandle.ResponseConfirm(roomId, mes)
	case pb.PbMessage_roomInit:
		messageHandle.roomMessageHandle.ResponseRoomInit(roomId, mes)
	case pb.PbMessage_select:
		messageHandle.roomMessageHandle.ResponseSelect(roomId, mes)
	case pb.PbMessage_load:
		messageHandle.roomMessageHandle.ResponseLoadResource(roomId, mes)
	case pb.PbMessage_fightStart:
		messageHandle.roomMessageHandle.ResponseFightStart(roomId, mes)

	}
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
	workerID := request.GetSession().GetSid() % 10
	messageHandle.TaskQueue[workerID] <- request

}

func NewMessageHandler(_server face.IServer) face.IMessageHandle {
	return &MessageHandle{
		server:         _server,
		WorkerPoolSize: 10,
		TaskQueue:      make([]chan face.IRequest, 10),
		// roomMessageHandle: &room.RoomMessageHandle{},
	}

}
