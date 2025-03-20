package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	// 返回对象
	return client
}

func main() {
	// 如果同一个目录下有两个main，只能分别build才不会引起冲突。
	// 否则最好分开成两个目录
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>连接服务器失败...")
		return
	}

	fmt.Println(">>>>连接服务器成功...")

	// 启动客户端的业务
	select {}
}
