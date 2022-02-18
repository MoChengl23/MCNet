package mnet

import (
	"encoding/binary"
	"fmt"
	"net"
	"server/face"
	match "server/mnet/Match"

	"sync"

	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	IP         string
	UDPIP      string
	sessionMap map[uint32]face.ISession //存放所有玩家的连接

	sessionMux sync.Mutex
	roomMap    map[uint32]face.IRoom //存放所有房间

	NewSessionSid uint32

	workerPool face.IWorkerPool
	matchSystem   face.IMatchSystem
}

func (server *Server) Start() {
	fmt.Println("Start Server")

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
		server.AddSession(conn)

	}
}

func (server *Server) Init() {
	server.matchSystem = match.NewMatchSystem(server)
	server.matchSystem.Init()

	server.workerPool = NewWorkerPool(server)
	server.workerPool.Init()
}

func (server *Server) Serve() {
	server.Init()
	server.Start()

	select {}

}
func (server *Server) AddSession(conn *kcp.UDPSession) {
	sid := conn.GetConv()

	newSession := NewSession(server, conn, sid)
	server.sessionMap[sid] = newSession
	newSession.Start()

	fmt.Println("a new session connect ")
}
func (server *Server) RemoveSession(sid uint32) {
	delete(server.sessionMap, sid)
}
func (server *Server) GetSession(sid uint32) face.ISession {
	if session, ok := server.sessionMap[sid]; ok {
		return session
	}
	return nil

}

func (server *Server) Stop() {

}
func (server *Server) GenerateUniqueSessionID() uint32 {
	server.sessionMux.Lock()
	for {
		_, ok := server.sessionMap[server.NewSessionSid]
		if !ok {
			break
		}
		server.NewSessionSid++
	}
	server.sessionMap[server.NewSessionSid] = nil

	server.sessionMux.Unlock()
	return server.NewSessionSid
}

func (server *Server) SendMessageToClient(sid uint32, data []byte) {
 
	server.sessionMap[sid].SendMessage(data)

}
func (server *Server) HandleMessage(request face.IRequest) {
	server.workerPool.AddToTaskQueue(request)
}

func (server *Server) GetMatchSystem() face.IMatchSystem {
	return server.matchSystem
}

func (server *Server) GetRoom(roomId uint32) face.IRoom {
	if room, ok := server.roomMap[roomId]; ok {
		return room
	}
	return nil
}
func (server *Server) AddRoom(roomId uint32, room face.IRoom) {
	server.roomMap[roomId] = room
}
func (server *Server) RemoveRoom(roomId uint32) {
	delete(server.roomMap, roomId)
}

//测试用
func (server *Server) GetAllPlayer() map[uint32]face.ISession {
	return server.sessionMap
}
func InitServerServive(server face.IServer) {

}

func NewServer() face.IServer {
	server := &Server{
		IP:    "0.0.0.0:6666",
		UDPIP: "0.0.0.0:7777",
		// messageHandle: NewMessageHandler(),
		sessionMap: make(map[uint32]face.ISession),
		roomMap:    make(map[uint32]face.IRoom),

		NewSessionSid: 1,
	}

	return server
}
