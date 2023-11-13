package main

import (
	"fmt"
	"net"
)

type Server struct{
	Ip   string
	Port int
}

// create a server interface
func NewServer (ip string, port int) *Server{
	server := &Server{
		Ip: ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler (conn net.Conn){
	// TODO: this connection's bussiness
	fmt.Println("Connection Established")
}


// start server interface
func (this *Server) Start(){
	// socker listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port) )
	if err != nil{
		fmt.Println("net.Listen err: ", err)
		return
	}
	// close listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err!= nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// handler
		go this.Handler(conn)

	}
	


	
}