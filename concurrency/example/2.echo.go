package main
/*
案例一中，clock服务器每一个链接都会起一个goroutine，
本例中会创建一个echo服务器，这个服务在每个连接中会有多个goroutine
大多数echo服务仅仅会返回他们读取到的内容
 */
