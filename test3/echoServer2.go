package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":7")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("echo server up")
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

	// 读取数据
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	// 发送相同数据
	_, err = conn.Write(buf[:n])
	if err != nil {
		return
	}
}
