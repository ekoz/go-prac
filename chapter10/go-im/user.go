package main

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string // 当前用户的channel
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 每次创建 user，则启动监听
	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	// 当前用户上线，要将用户加入到 server OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, " join now.")
}

// 用户下线业务
func (this *User) Offline() {
	// 当前用户上线，要将用户加入到 server OnlineMap
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, " quit.")
}

// 给当前用户对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + " online.\n"
			// 告诉当前用户
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[0:7] == "rename|" {
		// 消息格式：rename|张三
		newName := msg[7:]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg(newName + " has been used\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg(newName + " rename sccuss\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听当前 User Channel 的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
