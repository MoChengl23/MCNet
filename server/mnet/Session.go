package mnet

import (
	"fmt"
	"net"

	"server/face"
	"time"

	"github.com/xtaci/kcp-go/v5"
)

type Session struct {
	sid        uint32
	inRoom     bool
	room       face.IRoom
	inGame     bool
	kcpSession *kcp.UDPSession
	address    string
	isAlive    chan bool
	isInGame   bool //如果玩家在对局中则只需要转发byte，否则要解码

	messageChan chan []byte

	messageHandle face.IMessageHandle
}

func NewSession(messageHandle face.IMessageHandle, conn *kcp.UDPSession, sid uint32) face.ISession {
	session := &Session{
		sid:           sid,
		room:          nil, //记录对局中， 该玩家属于哪个间
		kcpSession:    conn,
		inRoom:        false,
		inGame:        false,
		messageChan:   make(chan []byte),
		messageHandle: messageHandle,
		isAlive:       make(chan bool),
		isInGame:      false,
	}
	return session
}

func (session *Session) CheckAlive() {
	defer session.Stop()
	for {
		select {
		case <-session.isAlive:

		case <-time.After(time.Second * 999):
			fmt.Println("Session die ")
			return
		}
	}
}

func (session *Session) StartReader() {
	fmt.Println("Session Start Read")
	defer session.Stop()

	for {
		buf := make([]byte, 256)

		n, err := session.kcpSession.Read(buf)
		//只截取有效数据部分
		buf = buf[:n]

		if err != nil {
			fmt.Println("session read data failed!!")
			break
		}

		fmt.Println(buf)
		request := Request{
			message: buf,
			session: session,
			conn:    session.GetConnection(),
		}

		session.messageHandle.AddToTaskQueue(&request)

		session.isAlive <- true
	}

}

func (session *Session) SendMessage(data []byte) {
	session.messageChan <- data
}

func (session *Session) StartWriter() {
	for {
		select {
		case data := <-session.messageChan:
			if _, err := session.kcpSession.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		}
	}

}

func (session *Session) Start() {
	go session.StartReader()

	go session.StartWriter()
	go session.CheckAlive()
}

func (session *Session) Stop() {
	session.kcpSession.Close()
	fmt.Println("session :", session.GetSid(), "STOP")

	//通知从缓冲队列读数据的业务，该链接已经关闭
	session.isAlive <- true

	//关闭该链接全部管道
	close(session.isAlive)
}

func (session *Session) GetConnection() net.Conn {
	return session.kcpSession
}

func (session *Session) GetSid() uint32 {
	return session.sid
}
func (session *Session) SetRoom(room face.IRoom) {
	session.room = room
}
func (session *Session) GetRoom() face.IRoom {
	return session.room
}

func (session *Session) IsInRoom() bool {
	return session.inRoom
}

func (session *Session) GetRemoteAddress() string {
	return session.address
}
