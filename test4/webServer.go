package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func unimplemented(conn net.Conn) {
	var buf string

	buf = "HTTP/1.0 501 Method Not Implemented\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "Server: httpd/0.1.0\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "Content-Type: text/html\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "<HTML><HEAD><TITLE>Method Not Implemented\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "</TITLE></HEAD>\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "<BODY><P>HTTP request method not supported.\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "</BODY></HTML>\r\n"
	_, _ = conn.Write([]byte(buf))
}

func handleCGI(path string) ([]byte, error) {
	// Set environment variables
	env := os.Environ()

	// Create CGI process
	cmd := exec.Command(path)
	cmd.Env = env
	stdout, _ := cmd.StdoutPipe()

	// Start CGI process
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("HTTP/1.1 500 Internal Server Error\r\n\r\n%s", err)
	}

	// Read CGI stdout
	var out []byte
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		out = append(out, buf[:n]...)
	}

	// Wait for CGI process to exit
	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return out, nil
}

func accept_request_thread(conn net.Conn) {
	defer conn.Close()
	var i int

	buf := make([]byte, 1024)
	n, err := conn.Read(buf) // 从conn中读取客户端发送的数据内容
	if err != nil {
		fmt.Printf("客户端退出 err=%v\n", err)
		return
	}

	// 获取方法
	i = 0
	var method_bt strings.Builder
	for i < n && buf[i] != ' ' {
		method_bt.WriteByte(buf[i])
		i++
	}
	method := method_bt.String()

	if method != "GET" {
		unimplemented(conn)
		return
	}

	for i < n && buf[i] == ' ' {
		i++
	}

	var url_bt strings.Builder
	for i < n && buf[i] != ' ' {
		url_bt.WriteByte(buf[i])
		i++
	}
	url := url_bt.String()

	if method == "GET" {

		var path, query_string string
		j := strings.IndexAny(url, "?")
		if j != -1 {
			path = url[:j]
			if j+1 < len(url) {
				query_string = url[j+1:]
			}
		} else {
			path = url
		}

		fmt.Println(path + "请求已经创建")
		resp := execute(path, query_string)
		header(conn, "text/html", len(resp))
		_, err := conn.Write(resp)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func header(conn net.Conn, content_type string, length int) {
	var buf string

	buf = "HTTP/1.0 200 OK\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "Server: httpd/0.1.0\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "Content-Type: " + content_type + "\r\n"
	_, _ = conn.Write([]byte(buf))
	_, _ = fmt.Sscanf(buf, "Content-Length: %d\r\n", length)
	buf = "Content-Type: " + content_type + "\r\n"
	_, _ = conn.Write([]byte(buf))
	buf = "\r\n"
	_, _ = conn.Write([]byte(buf))
}

func execute(path string, query_string string) []byte {
	query_params := make(map[string]string)
	parse_query_params(query_string, query_params)
	// camera_id := query_params["camera_id"]

	// resp := make(map[string]interface{})
	// resp["camera_id"] = camera_id
	// resp["code"] = 200

	if "/" == path {
		file, err := os.ReadFile("index.html")
		if err != nil {
			log.Fatal(err)
		}
		return file
	} else if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".txt") {
		file, err := os.ReadFile(path[1:])
		if err != nil {
			log.Fatal(err)
		}
		return file
	} else if strings.HasSuffix(path, ".sh") {
		out, _ := handleCGI(path[1:])
		return out
	}

	return []byte("no such path")
}

func parse_query_params(query_string string, query_params map[string]string) {
	kvs := strings.Split(query_string, "&")
	if len(kvs) == 0 {
		return
	}

	for _, kv := range kvs {
		kv := strings.Split(kv, "=")
		if len(kv) != 2 {
			continue
		}
		query_params[kv[0]] = kv[1]
	}
}
func main() {

	listen, err := net.Listen("tcp", ":8080") // 创建用于监听的 socket
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	fmt.Println("监听套接字，创建成功, 服务器开始监听。。。")
	defer listen.Close() // 服务器结束前关闭 listener

	// 循环等待客户端链接
	for {
		fmt.Println("等待客户端链接...")
		conn, err := listen.Accept() // 创建用户数据通信的socket
		if err != nil {
			panic("Accept() err=  " + err.Error())
		}
		// 这里准备起一个协程，为客户端服务
		go accept_request_thread(conn)
	}
}
