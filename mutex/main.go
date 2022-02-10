package main

import (
	"fmt"
	"time"
)

func main(){
	////基本用法，可以让各个goroutine互斥的访问临界资源x
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)

	////读多写少时
	//start := time.Now()
	//for i:=0;i<10;i++{
	//	go write()
	//	wg.Add(1)
	//}
	//for i:=0;i<1000;i++{
	//	go read()
	//	wg.Add(1)
	//}
	//wg.Wait()
	//fmt.Println(time.Since(start))

	//读写锁应用于读多写少时
	start := time.Now()
	for i:=0;i<10;i++{
		go RWwrite()
		wg.Add(1)
	}
	for i:=0;i<1000;i++{
		go RWread()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
