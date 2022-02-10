package main

import "net/rpc"

/*
rpc应用中，开发人员的三类角色：
1.服务端实现rpc方法的开发人员；
2.客户端调用rpc方法的人；
3.最重要：制定服务端与客户端rpc接口规范的设计人员
这样做便于后期维护
 */

//确定服务服务的名字和接口
const SkillServiceName = "path/to/pkg.HelloService"

type SkillServiceInterface = interface{
	Sing(req string,rep *string)error
	Swim(req string,rep *string)error
}

func RegisterSkillService(svc SkillServiceInterface)error{
	return rpc.RegisterName(SkillServiceName,svc)
}


