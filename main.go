package main

func main() {
	// TODO: Implement the main function
	// go mod init + GitHub项目地址 解决找不到newserver问题
	// 运行：go run server.go mian.go 启动服务端，然后再postman中向localhost:8888发送请求, 成功
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
