package main

import (
	"fmt"
	"time"
)

/*
goroutine是官方实现的超级”线程池“
go高并发的根本原因：每个实例只占栈内存的4~5KB，所以创建和销毁的开销是很小的
goroutine奉行通过通信来共享内存，而不是共享内存来通信（提倡使用channel，而不是lock）
goroutine由Go的运行时（runtime）调度和管理，省去了开发人员自己维护线程池的烦恼，也不必自己调度线程执行任务并维护上下文切换
Go程序会智能的将goroutine中的任务合理的分配给每个CPU，语言层面已经内置了调度和上下文切换的机制
注意：
程序启动时，Go会为main()函数创建一个主goroutine，当main()函数返回的时候，主goroutine就结束了，所有在main中启动的goroutine会一同结束
新的goroutine用go语句创建
 */

//开启两个goroutine，可以看到是并发执行的，打印顺序不固定
func start2goroutine(){
	go func() {
		for i :=0;i<5;i++{
			fmt.Println("this is goroutine1")
			time.Sleep(time.Second)
		}
		Wg.Done()
	}()
	go func() {
		for j:=0;j<5;j++{
			fmt.Println("this is goroutine2")
			time.Sleep(time.Second)
		}
		Wg.Done()
	}()
}

/*
操作系统线程（OS线程）一般都有固定的栈内存（通常为2MB）,一个个goroutine的栈在其生命周期开
始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，
goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大，所以在Go语言中一次创建十万左右的
goroutine也是可以的
 */