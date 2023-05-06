package main

import (
	"fmt"
	"net"
)

// 数据包方式的daytimeclient
func main() {
	var ip string
	fmt.Print("Enter target ip: ")
	fmt.Scanln(&ip)
	addr, err := net.ResolveUDPAddr("udp4", ip+":13")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 连接服务器
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 发送请求
	_, err = conn.Write([]byte("request"))
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

	fmt.Println(string(buf[:n]))
}
