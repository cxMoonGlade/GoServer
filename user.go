package main

import (
	"net"
)

type User struct{
	Name string
	Addr string
	C 	 chan string
	conn net.Conn
}

// Interface: create User, Passing in the connection 
func NewUser (conn net.Conn) *User{
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make (chan string),
		conn: conn,
	}
	// start the monitor
	go user.ListenMsg()
	return user
}

// Listen current User Channel
// once there is a msg, send to paired client
func (this *User) ListenMsg(){
	for {
		msg := <- this.C // will be blocked until there is a msg 

		// send the msg to client
		this.conn.Write([]byte(msg + "\n"))
	}
}

