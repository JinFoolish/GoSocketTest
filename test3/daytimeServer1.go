package main

import (
	"fmt"
	"net"
	"time"
)

// 数据包方式的daytimeserver
func main() {
	// 解析地址
	addr, err := net.ResolveUDPAddr("udp4", ":13")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 监听端口
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("daytime server up")
	defer conn.Close()

	for {
		// 接收客户端请求
		buf := make([]byte, 32)
		_, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 发送当前时间
		data := time.Now().Format("2006-01-02T15:04:05Z07:00")
		conn.WriteToUDP([]byte(data), clientAddr)
	}
}
