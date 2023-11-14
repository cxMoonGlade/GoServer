package main

import (
	"fmt"
	"net"
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

func main(){
	client:= NewClient("127.0.0.1", 8888)
	if client == nil{
		fmt.Println((">>>>> Connection to Server Failed..."))
		return
	}
	fmt.Println(">>>> Conncection Establish...")

	// TODO: start Client Business
	select{}
}