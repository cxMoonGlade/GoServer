# Go Chat Room

This is a Chat Room Server and Client written in Go. It provides a platform for users to connect, chat in public or private rooms, change their usernames, and stay connected in real-time.

## How to Run
```bash
go build -o server server.go user.go main.go
go build -o client client.go
```

then use
```bash
./server
```


to run the server instance, and run
```bash
./client
```
in different terminal windows to run multiple client instances.

You can also do 
```bash
./client -h
```
to check the manual, or
'''bash
./client -ip <Your Objected Ip Address> -port <Your Objected TCP Port Number>
'''
-ip means set the ip address of the server

-port means set the port number of the server


## Features

### Server

- **User Online and Broadcasting**: The server keeps track of online users and broadcasts notifications when users join or leave the chat room.

- **User Message Broadcasting**: Messages sent by users in the chat room are broadcasted to all other users in real-time.

- **Online User Querying**: Users can query the list of online users to see who is currently active in the chat room.

- **Inter-user Chatting**: Users can initiate one-on-one private chats with other users for private conversations.

- **User Name Changing**: Users have the ability to change their usernames.

- **Auto Kick-out Over Time**: Users who have been inactive for a specified amount of time can be automatically kicked out of the chat room.

### Client

- **Establish Connections**: Clients can establish connections with the server to join the chat room.

- **Public Chat**: Clients can participate in public chat rooms, where messages are visible to all connected users.

- **Private Chat**: Clients can initiate private conversations with other users for more personalized interactions.

- **User Name Changing**: Users have the ability to change their usernames.

- **Command Line Parsing**: Program can parse command line input to satisfy different needs

## Getting Started

### Prerequisites

- Go (Golang) must be installed on your system.

## Usage

- Start the server and multiple client instances to simulate a chat room environment.

- Clients can use the provided commands to perform various actions, such as joining rooms, sending messages, and initiating private chats.

- Experiment with different features of the chat room to understand how it works.

## Configuration

You can customize the behavior of the chat room server and client by modifying the configuration files provided in their respective directories.

## Contributing

Contributions are welcome! If you have any improvements or suggestions, please open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to the Go community for their excellent libraries and resources that make projects like this possible.

Feel free to add more details, installation instructions, usage examples, and other relevant information to make your README more comprehensive. Good luck with your project!
