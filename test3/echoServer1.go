package main

import (
	"fmt"
	"net"
)

func main() {
	// 解析地址
	addr, err := net.ResolveUDPAddr("udp4", ":7")
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
	fmt.Println("echo server up")
	defer conn.Close()

	for {
		// 读取数据
		buf := make([]byte, 32)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 发送相同数据
		_, err = conn.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
