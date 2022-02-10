package main

import "fmt"

/*
slice比起数组更加灵活，且功能强悍，也称为“动态数组”；
可以把切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，一定会包含一个数组；
切片也可以被看作是对数组的某个连续片段的引用，所以它是引用类型
 */

/*
Java 等编程语言中有令人困惑的“传值或传引用”问题；
在 Go 语言中，我们判断所谓的“传值”或者“传引用”只要看被传递的值的类型就好了。
如果传递的值是引用类型的，那么就是“传引用”；如果传递的值是值类型的，那么就是“传值”。
从传递成本的角度讲，引用类型的值往往要比值类型的值低很多。
 */

func base(){
	/*
	slice有两个概念：长度len和容量cap；
	len是指已经被赋过值的最大下标+1，可通过len(slice)获取；
	cap是指切片目前可容纳的最多元素个数，可通过cap(slice)获取
	 */
	slice0 := []int{1,2,3}	//1.直接初始化
	fmt.Println(len(slice0),cap(slice0))	//3 3

	/*
	slice是引用类型，传递slice时将引用同一指针，修改值将会影响其他对象
	 */
	array := [8]int{1,2,3,4,5,6,7,8}
	slice1 := array[:3]		//2.用数组初始化
	fmt.Println(len(slice1),cap(slice1))	//3 8
	slice1[0] = 6
	fmt.Println(slice1,array)	//[6 2 3] [6 2 3 4 5 6 7 8]

	/*
	除了直接初始化，用数组初始化，
	slice也可通过make初始化
	 */
	slice2 := make([]int,3)		//3.用make初始化
	slice3 := make([]int,2,3)
	fmt.Println(len(slice2),cap(slice2),len(slice3),cap(slice3))	//3 3 2 3

	/*
	两个切片指向同一个底层数组，分别修改两个切片的值，会有什么影响呢？
	 */
	slice4 := array[1:5]
	slice5 := array[1:6]
	slice4[0] = 100
	slice5[1] = 200
	fmt.Println(slice4 ,slice5)	//[100 200 4 5] [100 200 4 5 6]
}

//浅拷贝
func shallowcopy(s []int) []int {
	s[0] = 100
	return s
}


