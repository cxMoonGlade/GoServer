package main

import (
	"net"
	"strings"
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

	}else if  len(msg) > 7 && msg[:7] == "rename|"{
		// msg format: rename|Alex, rename is[0], Alex is[1]
		newName := strings.Split(msg, "|")[1]
		// determine if name is already exist
		_, ok := this.server.OnlineMap[newName]
		if ok{
			this.SendMsg("Sorry, name is already taken.\n")
		}else{
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("you has updated your name: " + this.Name + "\n")
		}
	}else if len(msg) > 4 && msg[:3] == "to|" {
		// msg format: to|name|msg content
		
		// 1. get the username
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == ""{
			this.SendMsg("message formate is incorrect.\nYou can try:to|Alex|Hello\n")
			return
		}
		// 2. get user object from username
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok{
			this.SendMsg("This user is not exist!\n")
			return
		}

		// 3. get msg content, send msg to the receiver
		content := strings.Split(msg, "|")[2]
		if content == ""{
			remoteUser.SendMsg(this.Name + "is very speechless for you\n")
		}else{
			remoteUser.SendMsg(this.Name+" said to you: " + content + "\n")
		}


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

