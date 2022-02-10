package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	listener,err := net.Listen("tcp","localhost:21999")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("server is listening...")
	for {
		conn,err := listener.Accept()	//Accept方法会直接阻塞，直到一个新的连接被创建
		if err != nil{
			log.Println(err)	// e.g. connection aborted
			continue
		}
		/*函数中的死循环会一直执行，直到写入失败
		最有可能的原因是客户端主动断开连接，这种情况下sendTime函数会用defer调用关闭服务器侧的连接；
		然后返回到主函数，继续等待下一个连接请求
		 */
		/*
		若有第二个客户端请求连接，则他必须等待第一个客户端完成工作
		因为此时的服务器程序同一时间只能处理一个客户端连接
		*/
		//sendTime(conn)
		/*
		此时的服务器就可以同时处理多个连接请求了
		 */
		go sendTime(conn)	//改为并发处理，让每一次sendTime调用都进入一个独立的goroutine中执行
	}
}
