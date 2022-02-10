package main

import (
	"fmt"
	"net"
)

func main(){
	//开始监听
	listen ,err := net.ListenUDP("udp",&net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 31000,
	})
	if err != nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	fmt.Println("listening...")
	defer listen.Close()
	for {
		var data [1024]byte
		n,addr,err := listen.ReadFromUDP(data[:])	//接收数据
		if err != nil{
			fmt.Println("read udp failed,err:",err)
			continue
		}
		fmt.Sprintf("data:%v addr:%v count:%v\n",string(data[:n]),addr,n)
		_,err = listen.WriteToUDP(data[:n],addr)	//写回客户端
		if err != nil{
			fmt.Println("write to udp failed,err:",err)
			continue
		}
	}
}
