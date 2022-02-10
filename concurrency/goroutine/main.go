package main

import (
	"sync"
	"time"
)
var Wg sync.WaitGroup
func main(){
	//Wg.Add(2)
	//start2goroutine()
	//Wg.Wait()

	//子goroutine执行到一半就退出
	go say()
	time.Sleep(1000*time.Second)

}
