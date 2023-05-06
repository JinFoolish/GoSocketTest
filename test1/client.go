package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090") //建立网络连接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//获取服务端消息
	go getMsg(conn)
	//将用户输入的文本消息发送到到服务端
	ioCopy(conn, os.Stdin)
}

func ioCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func getMsg(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		if n > 0 {
			if strings.Count(string(buf[0:n]), "exit") == 1 {
				fmt.Print("拜拜")
				conn.Close()
				os.Exit(0) //退出
			} else if string(buf[:4]) == "msg/" {
				fmt.Print(string(buf[4:n]))
			}
		}
	}
}
