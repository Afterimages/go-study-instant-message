package main

func main() {
	// TODO: Implement the main function
	// go mod init + GitHub项目地址 解决找不到newserver问题
	// 运行：go run server.go main.go 启动服务端，然后再postman中向localhost:8888发送请求, 成功
	server := NewServer("127.0.0.1", 8888)
	server.Start()
	// go build -o server.exe main.go server.go user.go
	// server
	// 或 go run server.go main.go user.go
	// 然后新开一个终端 chcp 65001 换行 ncat localhost 8888
}
