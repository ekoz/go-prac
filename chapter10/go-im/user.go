package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string // 当前用户的channel
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 每次创建 user，则启动监听
	go user.ListenMessage()

	return user
}

// 监听当前 User Channel 的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
