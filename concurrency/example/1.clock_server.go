package main

import (
	"io"
	"net"
	"time"
)

/*
这个案例中将完成一个顺序执行的时钟服务器，每隔一秒钟将当前的时间写到客户端
 */

/*
客户端使用nc工具，它可以用来执行网络连接操作:
nc localhost 21999
 */


func sendTime(c net.Conn){
	defer c.Close()
	for {
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil{
			return
		}
		time.Sleep(1 * time.Second)
	}
}
