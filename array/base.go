package main

import (
	"fmt"
	"reflect"
)

/*
数组是一组同类型数据的集合，是值类型，
通过从0开始的下标索引访问元素值
 */

func deepcopy(arr [3]int) [3]int{
	arr[0] = 4
	return arr
}

func changeValue(arr *[3]int){
	arr[0] = 4
}

func base(){
	/*
	数组在初始化之后，无法修改长度；
	数组长度也是其类型的一部分，所以[1]int和[2]int是不一样的类型，无法比较；
	数组长度声明时必须给定，通过内置函数len(array)获取；
	 */
	array1 := [3]int{1,2,3}
	fmt.Println("type of array1 is",reflect.TypeOf(array1),",length of it is",len(array1))

	/*
	数组是值类型：
	*将一个数组赋值给另外一个数组，实际上是将整个数组拷贝一份；
	*如果数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针；
	*修改数组的值需要传递数组的指针
	 */
	array2 := array1
	array2[0] = 4
	fmt.Println(array1,array2)
	fmt.Println(deepcopy(array1),array1)
	changeValue(&array1)
	fmt.Println(array1)


}