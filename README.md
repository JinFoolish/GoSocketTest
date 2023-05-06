# GoSocketTest

分布式系统实验

## 实验一  数据包socket应用

1. 构建客户端程序
（1） 构建datagramSocket对象实例
（2） 构建DatagramPacket对象实例，并包含接收者主机地址、接收端口号等信息
（3） 调用datagramSocket对象实例的send方法，将DatagramPacket对象实例作为参数发送。
2. 构建服务器端程序
（1） 构建datagramSocket对象实例，指定接收的端口号。
（2） 构建DatagramPacket对象实例，用于重组接收到的消息。
（3）调用datagramSocket对象实例大家receive方法，进行消息接收，并将DatagramPacket对象实例作为参数。  
3. 实现简单的聊天软件的功能
（1） 显示信息发送者的呢称（或IP地址），信息发送时间，信息内容。
（2） 实现多人聊天。

## 实验二  流式socket应用

1. 构建客户端程序和服务器端程序都需要的MystreamSocket类
2. 构建客户端程序
（1） 创建一个MyStreamsocket的实例对象，并将其指定接收服务器和端口号
（2） 调用该socket的receiveMessage方法读取从服务器端获得的消息

3. 构建服务器端程序
（1） 构建连接socket实例，并与指定的端口号绑定，该连接socket随时侦听客户端的连接请求
（2） 创建一个MyStreamsocket的实例对象
（3） 调用MyStreamsocket的实例对象的sendMessage方法，进行消息反馈。
4. 在实验一的聊天程序里添加发送图片和文件的功能
（1） 实现发送图片并显示图片功能。
（2） 实现发送文件并保存到指定位置功能。

## 实验三  客户/服务器应用开发

1. 数据包socket和流式socket实现daytime协议，从远程计算机上获取时间，并更新本机时间。
2. 数据包socket和流式socket实现echo协议，向远程计算机发送信息，并获取返回信息。

## 实验四  实现一个基本的web服务器程序

1. 创建 ServerSocket 类对象，监听端口 8080。这是为了区别于 HTTP 的标准 TCP/IP端口 80 而取的;
2. 等待、接受客户机连接到端口 8080，得到与客户机连接的 socket;
3. 创建与 socket 字相关联的输入流和输出流;
4. 从与 socket 关联的输入流 instream 中读取一行客户机提交的请求信息，请求信息的格式为：GET 路径/文件名 HTTP/1.0
5. 从请求信息中获取请求类型。如果请求类型是 GET，则从请求信息中获取所访问的文件名。没有 HTML 文件名时，则以 index.html 作为文件名;
6. 如果请求文件是 CGI 程序存则调用它，并把结果通过 socket 传回给 Web 浏览器，（此处只能是静态的 CGI 程序，因为本设计不涉及传递环境变量）然后关闭文件。否则发送错误信息给 Web 浏览器;
7. 关闭与相应 Web 浏览器连接的 socket 字。
