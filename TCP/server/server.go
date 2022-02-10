package main

import (
	"fmt"
	"net"
)

func main(){
	//利用 net 包实现tcp监听
	listen,err := net.Listen("tcp","127.0.0.1:21999")
	if err != nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	// 保持监听状态，循环接受请求并建立连接
	for {
		conn,err := listen.Accept() //建立连接
		if err != nil{
			fmt.Println("connect failed,err:",err)
			continue
		}
		go handleConn(conn) //启动一个goroutine处理连接
	}
}
