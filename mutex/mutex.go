/*
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
Go语言中使用sync包的Mutex类型来实现互斥锁。
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var wg sync.WaitGroup
var x int64
func add(){
	for i := 0;i<5000;i++{
		/*
		使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
		当互斥锁释放后，等待的goroutine才可以获取锁进入临界区；
		多个goroutine同时等待一个锁时，唤醒的策略是随机的。
		 */
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

/*
模拟读多写少的一个场景
 */
func read(){
	defer wg.Done()
	/*
	互斥锁在读的时候也会阻塞下一个读线程，效率比较低
	 */
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)	//读取需要1毫秒
	lock.Unlock()
}

func write(){
	defer wg.Done()
	lock.Lock()
	x ++
	time.Sleep(time.Millisecond*5)	//写入需要5毫秒
	lock.Unlock()
}