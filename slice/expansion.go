package main

import "fmt"

/*
一旦一个切片无法容纳更多的元素，Go 语言就会想办法扩容。
但它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中。
一般的情况下，可以简单地认为新切片的容量将会是原切片容量（以下简称原容量）的 2 倍。
 */

func expansion(){
	s := []int{1,2,3,4,5,6,7,8}
	s1 := s[3:6]
	fmt.Println("before expansion...")
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	//扩容
	fmt.Println("start expansion...")
	for i :=0;i<5;i++{
		s1 = append(s1,i)
	}
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	//第二次扩容，一次性扩充超过容量的数量
	fmt.Println("second expansion...")
	for i :=0;i<20;i++{
		s1 = append(s1,i)
	}
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
}

/*
特殊情况，当原切片的长度大于或等于1024时，Go 语言将会以原容量的1.25倍作为新容量的基准。
新容量基准会被调整（乘一次1.25，容量可能还是不够，所以需要不断地与1.25相乘），直到结果不小于原长度与要追加的元素数量之和。
最终，新容量往往会比新长度大一些，当然，相等也是可能的。
*/
func exp1024(){
	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
	fmt.Println()
}

/*
Q：	切片的底层数组什么时候会被替换？

A：	确切地说，一个切片的底层数组永远不会被替换。
	为什么？虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片。
	它只是把新的切片作为了新底层数组的窗口，而没有对原切片，及其底层数组做任何改动。
	请记住，在无需扩容时，append函数返回的是指向原底层数组的原切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。
	所以，严格来讲，“扩容”这个词用在这里虽然形象但并不合适。不过鉴于这种称呼已经用得很广泛了，我们也没必要另找新词了。
 */

/*
Q:	如果有多个切片指向了同一个底层数组，你认为应该注意些什么？
A：	两个方面考虑：
	1.初始时，两个切片引用同一个底层数组，后续操作中，如果对某个切片的操作超出底层数组的容量，这两个切片引用的就不是同一个数组了；
	2.当两个长度不一的切片使用同一个底层数组，并且两切片的长度均小于数组的容量时，对其中任何一个切片元素进行修改，都会影响另外一个切片值以及底层数组。
 */

/*
Q:	是否可以使用new来初始化slice？
A:	不能。make 专门用来创建 切片、map、channel 的值。它返回的是被创建的值，并且立即可用。
	new 是申请一小块内存并标记它是用来存放某个值的。它返回的是指向这块内存的指针，而且这块内存并不会被初始化。
	或者说，对于一个引用类型的值，那块内存虽然已经有了，但还没法用（因为里面没有针对那个值的数据结构的初始化操作）。
	所以，对于引用类型的值，不要用 new，能用 make 就用 make，不能用 make 就用复合字面量来创建。
	new 和make的区别：看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

	new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}
	make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。换言之，new 函数分配内存，make 函数初始化内存；
	slice、map以及channel都是golang内建的引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而make就是对三者进行初始化的一种操作方式
	new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以slice、map、channel需要make进行初始化并获取对应的内存地址，而非new简单的获取内存地址
 */
