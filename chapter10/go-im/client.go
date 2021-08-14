package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(ip string, port int) *Client {
	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
	}

	// 建立连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	return client
}

var serverIp string
var serverPort int

// debug: go run ./client.go -ip 127.0.0.1 -port 8080
// ./client -ip 127.0.0.1 -port 8080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip")
	flag.IntVar(&serverPort, "port", 8080, "设置服务器端口")
}

func main() {
	// 只有加上 flag.Parse()，才能真正的解析到命令行参数，否则一直都是默认值
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>> connect fail..")
		return
	}
	fmt.Println(">>>>>> connect seccuss..")

	// 启动客户端
	select {}
}
