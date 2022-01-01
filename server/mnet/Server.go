package mnet

import (
	"encoding/binary"
	"fmt"
	"net"
	"server/face"
	match "server/mnet/Match"
	"server/mnet/room"
	"sync"

	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	IP         string
	UDPIP      string
	sessionMap map[uint32]face.ISession //存放所有玩家的连接

	sessionMux sync.Mutex
	roomMap    map[uint32]face.IRoom //存放所有房间

	roomNumber   uint32 //当前共有多少个房间
	playerNumber uint32

	messageHandle face.IMessageHandle
	matchSystem   face.IMatchSystem
	// roomSystem  face.IRoomSystem
}

func (server *Server) Start() {
	fmt.Println("Start Server")
	//启动工作池
	// server.messageHandle.Init()
	go server.ListenUDP()

	go server.ListenKCP()
}

//UDP用于用户初次连接，分配一个KCPsession给他
func (server *Server) ListenUDP() {

	address, _ := net.ResolveUDPAddr("udp", server.UDPIP)
	// remoteaddress, _ := net.ResolveUDPAddr("udp", "127.0.0.1:7777")

	conn, err := net.ListenUDP("udp", address)

	if err != nil {
		fmt.Println("listen UDP failed!!")
	}
	for {

		buf := make([]byte, 32)
		_, remoteAddress, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("receive UDP failed!!")
		}
		// fmt.Println(remoteAddress)
		// fmt.Println(buf)
		//这里的连接是本地（只有一个）的连接吗？如果是的话还需要把udp连接关闭，暂时还不知道怎么写
		go func() {
			if binary.BigEndian.Uint32(buf[:4]) == 0 {

				//某个客户端申请连接,分配一个sid给它
				buf := make([]byte, 4)
				sid := server.GenerateUniqueSessionID()
				fmt.Println("UDPid", sid)
				binary.BigEndian.PutUint32(buf, sid)
				// fmt.Println("unique sid ", buf)
				// fmt.Println(append([]byte{0, 0, 0, 0}, []byte(buf)...))
				conn.WriteToUDP(append([]byte{0, 0, 0, 0}, []byte(buf)...), remoteAddress)

			}
		}()
	}
}

//KCP就是建立连接后的正常业务
func (server *Server) ListenKCP() {
	//开启服务器端口
	kcplisten, err := kcp.ListenWithOptions(server.IP, nil, 0, 0)
	if err != nil {
		fmt.Println("kcp.Listen failed!!")
	}
	for {
		//监听是否有新的客户端连接
		conn, err := kcplisten.AcceptKCP()
		if err != nil {
			fmt.Println("accept conn failed!!")
			continue
		}
		server.AddNewSession(conn)

		// for a, b := range server.sessionMap {
		// 	fmt.Println(a, b)
		// }

	}
}

func (server *Server) Serve() {
	server.Start()

	server.matchSystem = match.NewMatchSystem(server)
	server.matchSystem.Init()

	server.messageHandle = NewMessageHandler(server)
	server.messageHandle.Init()

	select {}

}
func (server *Server) AddNewSession(conn *kcp.UDPSession) {
	sid := conn.GetConv()
	if value, ok := server.sessionMap[sid]; ok {
		fmt.Println("sdfgsdf", value)
		delete(server.sessionMap, sid)

	}
	newSession := NewSession(server.messageHandle, conn, sid)
	server.sessionMap[sid] = newSession
	newSession.Start()

	fmt.Println("a new session connect ")
}

func (server *Server) Stop() {

}
func (server *Server) GenerateUniqueSessionID() uint32 {
	server.sessionMux.Lock()
	for {
		_, ok := server.sessionMap[server.playerNumber]
		if !ok {
			break
		}
		server.playerNumber++
	}
	server.sessionMap[server.playerNumber] = nil

	server.sessionMux.Unlock()
	return server.playerNumber
}

// func (server *Server) GenerateUniqueRoomID() int {
// 	server.roomMux.Lock()
// 	for {
// 		_, ok := server.roomMap[server.roomNumber]
// 		if !ok {
// 			break
// 		}
// 		server.roomNumber++
// 	}
// 	server.roomMux.Unlock()
// 	return server.roomNumber
// }

//输入是房间创造者的sid，房间号=房主的sid
func (server *Server) GenerateNewRoom(playerList [3]uint32) {
	newRoom := room.NewRoom(server, playerList)

	server.roomMap[playerList[0]] = newRoom
}

func (server *Server) SendMessageToClient(sid uint32, data []byte) {
	server.sessionMap[sid].SendMessage(data)
}

// server.roomMap[roomId] = room

// func (server *Server) LeaveRoom(roomId uint32, playerSid uint32) error {

// 	players := server.roomMap[roomId]

// 	for index, value := range players {
// 		if value == playerSid {
// 			players = append(players[:index], players[index+1:]...)
// 			break
// 		}
// 	}
// 	server.roomMap[roomId] = players
// 	if len(players) == 0 {
// 		server.RoomRemove(roomId)
// 	}
// 	//将该玩家的房间号重置
// 	err := server.sessionMap[playerSid].SetRoomId(0)
// 	return err
// }
func (server *Server) RoomRemove(roomId uint32) {

	delete(server.roomMap, roomId)

}
func (server *Server) GetSession(sid uint32) face.ISession {
	return server.sessionMap[sid]
}
func (server *Server) GetMatchSystem() face.IMatchSystem {
	return server.matchSystem
}

//开局，暂时先这么写
// func (server *Server) GameStart(roomId uint32, roomPlayers []uint32) error {
// 	//（保险起见）再次更新 此时房间内的人id
// 	server.roomMap[roomId] = roomPlayers
// 	for _, playerSid := range roomPlayers {
// 		server.sessionMap[playerSid].SetRoomId(roomId)
// 	}
// 	return nil

// }

// func (server *Server) RoomBroadCast(roomId uint32, message []byte) {
// 	for _, roomPlayers := range server.roomMap[roomId] {
// 		server.sessionMap[roomPlayers].SendMessage(message)
// 	}
// }

func NewServer() face.IServer {
	server := &Server{
		IP:    "0.0.0.0:6666",
		UDPIP: "0.0.0.0:7777",
		// messageHandle: NewMessageHandler(),
		sessionMap:   make(map[uint32]face.ISession),
		roomMap:      make(map[uint32]face.IRoom),
		roomNumber:   0,
		playerNumber: 1,
	}
	server.messageHandle = NewMessageHandler(server)
	return server
}

// func NewMessageHandler() *MessageHandle {
// 	return &MessageHandle{

// 		WorkerPoolSize:    10,
// 		TaskQueue:         make([]chan face.IRequest, 10),
// 		roomMessageHandle: &room.RoomMessageHandle{},
// 	}
// }

// func NewRoomMessageHandle() *RoomMessageHandle {
// 	return &RoomMessageHandle{}

// }
