package ee

import (
	"net"
	// "server/pb"
	// "google.golang.org/protobuf/proto"
)

type Session struct {
	sid uint32
	// roomid    uint32
	// kcpLister net.Listener //收
	kcpSession net.Conn //发
	raddr      string
	c          chan []byte
}

func (sess *Session) InitSession(sid uint32, raddr string, kcpSession net.Conn) {
	sess.raddr = raddr
	sess.sid = sid

	sess.kcpSession = kcpSession

}
func (sess *Session) ListenMessage() []byte {
	buf := make([]byte, 128)
	for {
		_, err := sess.kcpSession.Read(buf)
		if err != nil {

		}
		return buf

	}

}

//监听user的channel，如果有消息则发送给客户端
func (sess *Session) SendMessage() {
	for {
		msg := <-sess.c
		sess.kcpSession.Write(msg)
	}

}

// func (sess *Session) Serve() {
// 	for {
// 		// data := make([]byte, 128)
// 		// _, _, error := sess.conn.ReadFromUDP(data)
// 		// if error != nil {
// 		// 	fmt.Println(error)
// 		// 	continue
// 		// }

// 		// fmt.Println(string(data[20:n]))
// 		// fmt.Println(data)
// 		_, _ = sess.listener.Write([]byte(string("aaa")))
// 		time.Sleep(time.Duration(1) * time.Second)
// 		fmt.Println(sess.listener.GetConv())
// 		// _, error = sess.conn.WriteToUDP([]byte(string(data)), remoteAddress)

// 	}
// 	select {}
// }

// func NewSess() *Session {
// 	var Laddr = "0.0.0.0:6666"
// 	address, _ := net.ResolveUDPAddr("udp", Laddr)
// 	remoteaddress, _ := net.ResolveUDPAddr("udp", "127.0.0.1:7777")

// 	conn, _ := net.ListenUDP("udp", address)
// 	sess := new(Session)
// 	sess.conn = conn
// 	sess.sid = 3
// 	sess.listener, _ = kcp.NewConn3(3, remoteaddress, nil, 0, 0, conn)
// 	return sess

// }

// func main() {
// 	// m := make(map[string]int)
// 	// go listenUDP()

// 	kcpconn, err := kcp.Dial(3, "localhost:6666")
// 	kcpconn.Write([]byte("edfg"))
// 	// kcplisten, err := kcp.Listen("127.0.0.1:7777")
// 	address, _ := net.ResolveUDPAddr("udp", "127.0.0.1:7777")
// 	conn, _ := net.ListenUDP("udp", address)
// 	// go listenKcp(conn, kcpconn)

// 	if err != nil {
// 		panic(err)
// 	}

// go func() {
// 	// var buffer = make([]byte, 16)
// 	for {
// 		kcpconn.Write(data)
// 		time.Sleep(time.Duration(1) * time.Second)

// 	}

// }()

// 	select {}

// }

// func pbHandle() []byte {
// 	student := &pb.Student{
// 		Name:   "mocheng",
// 		Male:   true,
// 		Scores: []int32{1, 1, 1, 1, 2, 132, 3, 13, 123},
// 	}
// 	data, error := proto.Marshal(student)
// 	if error != nil {

// 	}
// 	return data
// }

// func GetPbMessage(bytes []byte) *pb.Student {
// 	message := &pb.Student{}
// 	error := proto.Unmarshal(bytes, message)
// 	if error != nil {

// 	}
// 	return message

// }

// func handleEcho(conn net.Conn) {
// 	fmt.Println("new client")
// 	buf := make([]byte, 32)
// 	for {
// 		n, err := conn.Read(buf)

// 		// fmt.Println("read!!")
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		fmt.Println(string(buf[:n]))
// 		fmt.Println(buf)
// 		n, err = conn.Write(buf[:n])
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// // package main

// // import (
// // 	"fmt"
// // 	"kcp-go"
// // 	"net"
// // )

// // func main() {

// // 	var Laddr = "0.0.0.0:6666"
// // 	address, _ := net.ResolveUDPAddr("udp", Laddr)
// // 	// remoteaddress, _ := net.ResolveUDPAddr("udp", "127.0.0.1:7777")

// // 	conn, _ := net.ListenUDP("udp", address)
// // 	go func() {
// // 		for {
// // 			data := make([]byte, 64)
// // 			_, _, error := conn.ReadFromUDP(data)
// // 			if error != nil {
// // 				fmt.Println(error)
// // 				continue
// // 			}

// // 			// fmt.Println(string(data[20:n]))
// // 			fmt.Println(data)
// // 		}
// // 	}()
// // 	// fmt.Println(remoteaddress)
// // 	// _, _ = sess.listener.Write([]byte(string("aaa")))
// // 	// time.Sleep(time.Duration(1) * time.Second)
// // 	// fmt.Println(listener.GetConv())
// // 	// _, error = sess.conn.WriteToUDP([]byte(string(data)), remoteAddress)

// // 	// lis, _ := kcp.ListenWithOptions("0.0.0.0:6666", nil, 10, 3)
// // 	// go func() {
// // 	// 	for {

// // 	// 		conn, _ := lis.AcceptKCP()
// // 	// 		var buffer = make([]byte, 64)
// // 	// 		n, _ := conn.Read(buffer)

// // 	// 		fmt.Println("receive from client:", buffer[:n])
// // 	// 		conn.Write([]byte("你好"))

// // 	// 		fmt.Println(buffer)
// // 	// 		time.Sleep(time.Duration(1) * time.Second)
// // 	// 	}
// // 	// }()
// // 	kcpconn, err := kcp.DialWithOptions("localhost:6666", nil, 10, 3)
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	go func() {
// // 		var buffer = make([]byte, 16)
// // 		for {
// // 			n, e := kcpconn.Read(buffer)
// // 			if e != nil {
// // 				fmt.Println(e.Error())
// // 				break
// // 			}
// // 			fmt.Println(string(buffer[:n]))
// // 		}

// // 	}()
// // 	kcpconn.Write([]byte("edfg"))

// // 	//time.Sleep(5 * time.Second)
// // 	// kcpconn.Write([]byte{7, 7, 7, 7})
// // 	// time.Sleep(time.Duration(1) * time.Second)

// // 	select {}

// // }
