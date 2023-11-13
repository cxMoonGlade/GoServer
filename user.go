package main

import (
	"net"
)

type User struct{
	Name string
	Addr string
	C 	 chan string
	conn net.Conn

	// belongs to which server
	server *Server
}

// Interface: create User, Passing in the connection 
func NewUser (conn net.Conn, server *Server) *User{
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make (chan string),
		conn: conn,
		server : server,
	}
	// start the monitor
	go user.ListenMsg()
	return user
}

func (this *User) Online(){
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// boradcasting user online msg
	this.server.BroadCast(this, "is ONLINE!")

}

func (this *User) Offline(){
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.BroadCast(this, "is OFFLINE")
}

// send msg to the client of the user
func (this *User) SendMsg(msg string){
	this.conn.Write([]byte(msg))
}

func (this *User) MessageHandler(msg string){
	// quering all online users
	if msg == "$OL"{
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap{
			omsg := "[" + user.Addr + "]" + user.Name + ":" + " is online...\n"
			this.SendMsg(omsg)
		}
		this.server.mapLock.Unlock()

	}else{
		this.server.BroadCast(this, msg)
	}
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

