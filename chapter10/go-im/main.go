package main

func main() {
	// go build -o server.exe main.go server.go
	// ./server.exe
	// telnet 127.0.0.1 8080
	server := NewServer("127.0.0.1", 8080)
	server.Start()
}
