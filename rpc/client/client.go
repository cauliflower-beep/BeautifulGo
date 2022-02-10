package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
const SkillServiceName = "path/to/pkg.HelloService"
func main(){
	//gob版客户端
	//client ,err := rpc.Dial("tcp","localhost:6672")
	//if err != nil{
	//	log.Fatal("dial err:",err)
	//}

	//json版客户端
	conn ,err := net.Dial("tcp","localhost:6672")
	if err != nil{
		log.Fatal("dial err:",err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var rep string
	//1.raw
	//err = client.Call("rawrpc.Hello","lufy!",&rep)
	//if err != nil{
	//	log.Fatal("rpc service err:",err)
	//}

	//2.normal
	err = client.Call(SkillServiceName+".Swim","lufy ",&rep)
	if err != nil{
		log.Fatal("rpc service err:",err)
	}
	fmt.Println(rep)
}
