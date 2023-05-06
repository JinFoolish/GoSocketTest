package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var ip string
	fmt.Print("Enter target ip: ")
	fmt.Scanln(&ip)
	// 解析服务器地址
	addr, err := net.ResolveTCPAddr("tcp4", ip+":13")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 连接服务器
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 发送请求
	_, err = conn.Write([]byte("1"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 每秒接收时间
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 读取时间
			buf := make([]byte, 32)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			// 打印时间
			data := string(buf[:n])
			fmt.Printf("%s\n", data)
		}
	}
}
