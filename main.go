package main

func main(){
	ip := "127.0.0.1"
	port := 8888
	server := NewServer(ip, port)
	server.Start()
}