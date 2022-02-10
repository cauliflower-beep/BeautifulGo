package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	//1.raw
	//rpc.RegisterName("rawrpc",new(RawRpc))

	//2.normal
	err := RegisterSkillService(new(SkillService))
	if err != nil{
		log.Fatal("register failed:",err)
	}

	//建立监听
	listener ,err := net.Listen("tcp",":6672")
	if err != nil{
		log.Fatal("listen tcp err:",err)
	}
	fmt.Println("listening establish...")
	//连接一次
	//conn,err := listener.Accept()
	//if err != nil{
	//	log.Fatal("listen tcp err:",err)
	//}
	//rpc.ServeConn(conn)

	//循环建立连接
	for{
		conn,err := listener.Accept()
		if err != nil{
			log.Fatal("listen tcp err:",err)
		}
		//默认gob编码的rpc服务
		//go rpc.ServeConn(conn)

		//为了方便跨语言调用，这里可以使用json编码
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

//2.normal:基于RPC接口规范编写的真实的服务端代码
type SkillService struct {}

func (s *SkillService)Sing(req string,rep *string)error{
	*rep = req + "can sing~"
	return nil
}

func (s *SkillService)Swim(req string,rep *string)error{
	*rep = req + "can swim~"
	return nil
}
