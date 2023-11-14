package main

import (
	"fmt"
	"net"
	"flag"
)

type Client struct{
	ServerIP string
	ServerPort int
	Name string
	conn net.Conn
	flag int // current mode
}

func NewClient(sip string, sport int) *Client{
	client := &Client{
		ServerIP:   sip,
		ServerPort: sport,
		flag:       9,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", sip, sport))
	if err != nil{
		fmt.Println("net.Dial error:",err)
		return nil
	}
	client.conn = conn
	return client
}

func (client *Client) menu() bool{
	var flag int
	fmt.Println("1. Broadcast mode")
	fmt.Println("2. Private Chat mode")
	fmt.Println("3. Update User Name")
	fmt.Println("0. Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3{
		client.flag = flag
		return true
	}else{
		fmt.Println(">>>>>Please Input Integer Number within Legal Range...")
		return false
	}
}

func (client *Client) Run(){
	for client.flag != 0{
		for client.menu() != true{

		}
		switch client.flag{
		case 1:
			// broadcast mode
			fmt.Println("Broadcasting Mode Selected...")
			break
		case 2:
			// private chat mode
			fmt.Println("Private Chat Mode Selected...")
			break
		case 3:
			// update user name
			fmt.Println("Update User Name...")
			break
		case 0:
			// Exit 
			fmt.Print("Thx for Using, Bye...")
			return
		}
	}
}


var sIP string
var sPort int

// ./client -ip 127.0.0.1

func init(){
	flag.StringVar(&sIP, "ip", "127.0.0.1", "set server ip address(default 127.0.0.1)")
	flag.IntVar(&sPort, "port", 8888, "set server tcp port(default 8888)")
}

func main(){
	// command line parse
	flag.Parse()

	client:= NewClient(sIP, sPort)
	if client == nil{
		fmt.Println((">>>>> Connection to Server Failed..."))
		return
	}
	fmt.Println(">>>> Conncection Establish...")

	// TODO: start Client Business
	client.Run()
}