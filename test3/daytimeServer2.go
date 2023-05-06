package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":13")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("daytime server up")
	for {
		// accept 连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 处理连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 读取请求
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	fmt.Printf("%s\n", string(buf[:1]))
	if err != nil {
		return
	}

	// 每秒发送时间
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data := time.Now().Format("2006-01-02T15:04:05Z07:00")
			conn.Write([]byte(data + "\n"))
		}
	}
}
