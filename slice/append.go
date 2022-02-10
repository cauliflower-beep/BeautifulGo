package main

import "fmt"

/*
1.	将切片 b 的元素追加到切片 a 之后：a = append(a, b…)

2.	赋值切片a的元素到新的切片b上：
	b = make([]T, len(a))
	copy(b, a)

3.	删除位于索引 i 的元素：a = append(a[:i], a[i+1:]…)

4.	切除切片 a 中从索引 i 至 j 位置的元素([i, j)：a = append(a[:i], a[j:]…)

5.	为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)…)

6.	在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]…)…)

7.	在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]…)…)

8.	在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]…)…)


9.	取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]

10.	将元素 x 追加到切片 a：a = append(a, x)
 */

//append添加的位置
func loc(){
	s1 := make([]int,0)		//初始化长度为0，则append的时候从第一位开始
	s2 := make([]int,3)		//初始化长度为3，则append的时候从第4位开始
	for i :=1;i<=5;i++{
		s1 = append(s1,i)	//[1 2 3 4 5]
		s2 = append(s2,i)	//[0 0 0 1 2 3 4 5]
	}
	fmt.Println(s1,"\n",s2)
}
