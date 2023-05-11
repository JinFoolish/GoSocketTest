package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 创建结构体存储用户名称、信息和时间
type chatMsg struct {
	name string
	msg  string
	t    string
}

var (
	msg      = make(chan string)  //用于新用户加入时广播信息
	chatMsgs = make(chan chatMsg) //发送聊天信息
	exit     = make(chan chatMsg) //离开信息
	trans    = make(chan chatMsg) //传输信息
)

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:10000") //监听本机9090端口
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server start")
	clients := make(map[string]net.Conn) //定义一个map，使用用户名作为key，value该用户的链接
	go func() {                          //开启一个线程，不断地从通道中读取数据
		for {
			select {
			case chat := <-chatMsgs: //有用户发送消息
				for name, accept := range clients {
					if name == chat.name { //当发送者是本人
						fmt.Fprint(accept, chat.t+" 你说："+chat.msg+"\n"+"=====>")
					} else {
						fmt.Fprint(accept, chat.t+" "+chat.name+" 说："+chat.msg+"\n"+"=====>")
					}

				}
			case exitMsg := <-exit: //有用户离开了
				for name, accept := range clients {
					if name == exitMsg.name {
						accept.Write([]byte("exit"))
					} else {
						fmt.Fprintln(accept, "<"+exitMsg.msg+">")
					}
				}
				delete(clients, exitMsg.name)
			case m := <-msg: //上线
				for _, accept := range clients {
					fmt.Fprint(accept, m)
				}
			}
		}
	}()

	for { //循环读取客户端接入
		accept, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("has client")
		go func() {
			client := accept.RemoteAddr().String()                                         //客户端地址
			name := "guest" + client[strings.LastIndex(client, ":")+1:len([]rune(client))] //使用端口号作为名称后缀
			clients[name] = accept
			fmt.Print(client) //将用户存入map
			msg <- "<" + name + "已经加入群聊" + ">" + "\n" + "======>"
			input := bufio.NewScanner(accept)
			for input.Scan() { //循环接收内容
				text := input.Text()
				if text == "exit" { //客户端输入"exit"消息时，退出聊天室
					exit <- chatMsg{name, name + "已退出", ""}
					break //跳出循环读取
				} else {
					t := time.Now()                                                    //获取当前时间
					chat := chatMsg{name, input.Text(), t.Format("02 Jan 2006 15:04")} //格式化时间
					chatMsgs <- chat
				}
			}
			if err := input.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input", err)
			}
		}()
	}

}
