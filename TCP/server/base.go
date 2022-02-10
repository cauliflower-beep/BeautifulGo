package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
TCP服务端程序的处理流程：
 1.监听端口
 2.接收客户端请求建立链接
 3.创建goroutine处理链接。
 */

//连接建立之后的处理函数
func handleConn(conn net.Conn){
	defer conn.Close()	//	关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])	//读取数据
		if err != nil{
			fmt.Println("read from client failed,err:",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：",recvStr)
		conn.Write([]byte(recvStr))	//写回客户端
	}
}
