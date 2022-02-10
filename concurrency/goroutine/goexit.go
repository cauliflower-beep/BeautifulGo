package main

import (
	"fmt"
	"runtime"
)

/*
如何让开启的goroutine跑一半就结束呢？
——在想要终止的地方插入"runtime.Goexit()"
 */

/*
只要执行 goexit 这个函数，当前协程就会退出，也就是说，P可以将当前g从M上摘下来，同时调度下一个可执行的协程g，放到M上去执行
 */

func say(){
	fmt.Println("11111")
	defer fmt.Println("22222")
	runtime.Goexit()	//执行到这里就结束了，结束前还能执行到defer
	fmt.Println("33333")
}

/*
goexit的用途有哪些呢？
 */

/*
一个goroutine可以打断另外一个goroutine的执行吗？
首先需要明确一点：除了从主函数退出或者直接终止程序之外，没有其他的编程方法能够让一个goroutine打断另一个的执行。
但是有某种方法可以实现这个目的：
通过goroutine之间的通信来让一个goroutine请求其他的goroutine，并被请求的goroutine自行结束执行
 */
