package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(ip string, port int) *Client {
	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
		flag:       999,
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

// 处理server端返回的消息，直接显示在控制台
func (client *Client) HandleResponse() {
	// 一旦 client 有数据，就直接copy到标准输出，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1-聊天室模式")
	fmt.Println("2-私人聊天模式")
	fmt.Println("3-修改名称")
	fmt.Println("0-退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法的序号<<<<<<")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 聊天室模式
func (client *Client) RoomChat() {
	var chatMsg string
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空，则发送至服务器
		if len(chatMsg) > 0 {
			chatMsg = chatMsg + "\n"
			_, err := client.conn.Write([]byte(chatMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) SelectUsers() {
	_, err := client.conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	client.SelectUsers()
	var receiver string
	var chatMsg string
	fmt.Println("请输入聊天对象用户名，exit退出")
	fmt.Scanln(&receiver)

	for receiver != "exit" {
		// 消息不为空，则发送至服务器
		if len(receiver) > 0 {
			fmt.Println("请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)

			for chatMsg != "exit" {
				// 消息不为空，则发送至服务器
				if len(chatMsg) > 0 {
					chatMsg = "to|" + receiver + "|" + chatMsg + "\n"
					_, err := client.conn.Write([]byte(chatMsg))
					if err != nil {
						fmt.Println("conn.Write err:", err)
						break
					}
				}

				chatMsg = ""
				fmt.Println("请输入聊天内容，exit退出")
				fmt.Scanln(&chatMsg)
			}
		}

		receiver = ""
		fmt.Println("请输入聊天对象用户名，exit退出")
		fmt.Scanln(&receiver)
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			fmt.Println("进入聊天室")
			client.RoomChat()
			break
		case 2:
			fmt.Println("私人聊天")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("修改名称")
			client.UpdateName()
			break
		}

	}
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

	// 单独开启一个 goroutine 处理服务器端的消息
	go client.HandleResponse()

	// 启动客户端
	client.Run()
}
