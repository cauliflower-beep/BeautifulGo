package main

import "fmt"

/*
1.map是一个无序的key/value对的集合，通过key可以在常数时间复杂度内检索、更新或删除对应的value；
2.map中所有的key都有相同的数据类型，所有value也都有相同类型，但key和value之间可以是不同类型；
这与python不同，python中的字典，key（value）可以是不同的数据类型；
3.map中的key必须是支持"=="比较运算符的数据类型，所以可以通过测试key是否相等来判断是否已存在；
4.虽然浮点类型支持"=="运算符比较，但最好不要将浮点数作为map的key类型，最坏情况下，可能出现任何key都不相等；
 */
func base(){
	//1.创建map
	//1.1内置的make创建
	ages := make(map[string]int)
	fmt.Println(ages)
	//1.2map字面值创建，同时初始化一些key/value
	addr := map[string]string{
		"alice":	"newyork",
		"kevin":	"wuhan",
	}
	fmt.Println(addr)
	/*
	1.2相当于：
	addr := make(map[string]int)
	addr["alice"] = "newyork"
	addr["charlie"] = "wuhan"
	*/
}
