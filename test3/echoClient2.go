package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务器
	var ip string
	fmt.Print("Enter target ip: ")
	fmt.Scanln(&ip)
	conn, err := net.Dial("tcp", ip+":7")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 接收响应
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Received: %s\n", string(buf[:n]))
}
