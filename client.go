package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // current mode
}

func NewClient(sip string, sport int) *Client {
	client := &Client{
		ServerIP:   sip,
		ServerPort: sport,
		flag:       9,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", sip, sport))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1. Broadcast mode")
	fmt.Println("2. Private Chat mode")
	fmt.Println("3. Update User Name")
	fmt.Println("0. Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>Please Input Integer Number within Legal Range...")
		return false
	}
}

// query online user
func (client *Client) SelectUser(){
	sendMSG := "$OL\n"
	_, err := client.conn.Write([]byte(sendMSG))
	if err != nil{
		fmt.Println("conn.Write error:", err)
		return
	}
}

func (client *Client) PrivateChat(){
	var remoteName string
	var chatMSG string

	client.SelectUser()
	fmt.Println(">>>>>Please enter who the message sent to[username]:exit to quit")
	fmt.Scanln(&remoteName)

	for remoteName != "exit"{
		fmt.Println(">>>>Please enter your MESSAGE:exit to quit")
		fmt.Scanln(&chatMSG)
		for chatMSG != "exit"{
		// if msg is not empty then send
			if len(chatMSG) != 0{
				sendMSG := "to|" + remoteName + "|" + chatMSG + "\n"
				_, err := client.conn.Write([]byte(sendMSG))
				if err != nil{
					fmt.Println("conn Write Error:", err)
					break
				}
			}
			chatMSG = ""
			fmt.Println(">>>>Please enter your MESSAGE:exit to quit")
			fmt.Scanln(&chatMSG)
		}
		remoteName = ""
		fmt.Println(">>>>>Please enter who the message sent to[username]:exit to quit")
		fmt.Scanln(&remoteName)
	}
}


func (client *Client) BroadCasting() {
	var chatMSG string
	fmt.Println(">>>>Please enter your msg, exit to quit")
	fmt.Scanln(&chatMSG)

	for chatMSG != "exit" {

		// if msg isn't empty
		if len(chatMSG) != 0 {
			sendMSG := chatMSG + "\n"
			_, err := client.conn.Write([]byte(sendMSG))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}

		chatMSG = ""
		fmt.Println(">>>>Please enter your msg, exit to quit")
		fmt.Scanln(&chatMSG)
	}
}

func (client *Client) UpdateUserName() bool {
	fmt.Println(">>>>>Input Your New User Name:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"

	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// hanlder msg from server, print to stdout
func (client *Client) ResponseHandler() {
	// if clent.conn has data, directly copy to std, permananently block and moniting
	io.Copy(os.Stdout, client.conn)

}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			// broadcast mode
			client.BroadCasting()
			break
		case 2:
			// private chat mode
			client.PrivateChat()
			break
		case 3:
			// update user name
			client.UpdateUserName()
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

func init() {
	flag.StringVar(&sIP, "ip", "127.0.0.1", "set server ip address(default 127.0.0.1)")
	flag.IntVar(&sPort, "port", 8888, "set server tcp port(default 8888)")
}

func main() {
	// command line parse
	flag.Parse()

	client := NewClient(sIP, sPort)
	if client == nil {
		fmt.Println((">>>>> Connection to Server Failed..."))
		return
	}

	go client.ResponseHandler()

	fmt.Println(">>>> Conncection Establish...")

	// TODO: start Client Business
	client.Run()
}
