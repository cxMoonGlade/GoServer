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
}

func NewClient(sip string, sport int) *Client{
	client:=  &Client{
		ServerIP: sip,
		ServerPort: sport,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", sip, sport))
	if err != nil{
		fmt.Println("net.Dial error:",err)
		return nil
	}
	client.conn = conn
	return client
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
	select{}
}