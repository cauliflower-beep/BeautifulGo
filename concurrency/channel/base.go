package main
/*
channel是goroutines之间的通信机制
一个goroutine可以通过channel给另一个goroutine发送值信息
channel也有类型，即他可以发送数据的类型，
例如一个可以发送int类型数据的channel，一般写为chan int
channel可以通过make函数来创建，是引用类型，零值为nil
 */

/*
无缓存channel：
一个基于无缓存channel的发送操作将导致发送者goroutine阻塞，
直到另一个goroutine在相同的channel上执行接受操作。
当发送的值通过channel成功传输之后，两个goroutine可以继续执行后面的语句。
反之，如果接受操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的channel上执行发送操作
 */


