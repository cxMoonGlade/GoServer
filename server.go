package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// current online user list
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// channel to broadcast
	Message chan string
}

// create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

// Goroutine to Listen Message broadcasting channel,
// once there is a msg, send to all online user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// sned msg to all online user
		this.mapLock.Lock()
		// OnlineMap : [Name string] User *User
		// Get the User Instance which is the value
		for _, cli := range this.OnlineMap {
			// send msg to User.C
			//notify each user one by one on the map
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}


func (this *Server) Handler(conn net.Conn) {
	// TODO: this connection's bussiness
	// connection established, get the user info

	user := NewUser(conn, this)

	user.Online()

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			// if there is an err and not because of we reached the end of file
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// abstract user msg(remove "\n")
			msg := string(buf[:n-1])

			// broadcast the msg
			user.MessageHandler(msg)
		}
	}()

	// current handler block
	select {}
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//start bc-msg monitor goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}
