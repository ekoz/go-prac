package main

import "fmt"

func main() {
	// debug : go run main.go user.go server.go
	// go build -o server.exe main.go server.go
	// ./server.exe
	// telnet 127.0.0.1 8080
	// Ctrl + ] and send message
	// win10 nc command https://joncraton.org/blog/46/netcat-for-windows/
	fmt.Println("welcome to go-im project")
	server := NewServer("127.0.0.1", 8080)
	server.Start()
}
