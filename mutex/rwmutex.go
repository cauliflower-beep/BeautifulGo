package main

import (
	"fmt"
	"sync"
	"time"
)

/*
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的;
当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的;
这种场景下使用读写锁(sync.RWMutex)是更好的一种选择。
 */

/*
读写锁分为两种：读锁和写锁。
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
 */

var rwlock sync.RWMutex
func RWread(){
	defer wg.Done()
	rwlock.RLock()		//加读锁，禁止其他线程写入，但不禁止读取
	fmt.Println(x)
	time.Sleep(time.Millisecond)	//读耗时1毫秒
	rwlock.RUnlock()	//解读锁
}

func RWwrite(){
	defer wg.Done()
	rwlock.Lock()		//加写锁，禁止其他线程读取或者写入
	x ++
	time.Sleep(time.Millisecond*5)		//写耗时5毫秒
	rwlock.Unlock()		//解写锁
}
