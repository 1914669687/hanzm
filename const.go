package main

import "fmt"

//常量，定义之后不能修改

func main() {
	const pi = 3.1415926
	const exp = 2.71828
	fmt.Println(pi, exp)

	const (
		n1 = 1
		n2
		n3 //批量声明常量时，如果有一行没有被赋值，则默认与上一行相同
	)
	fmt.Println(n1, n2, n3)

	const (
		m1 = iota
		m2
		m3 //iota是golang的常量计数器，只能在常量表达式中使用。
		//在const关键字出现时被置为0，const中每新增一行变量声明，iota计数一次
	)
	fmt.Println(m1, m2, m3)

	const (
		r2 = 10
		r3 = 11
		r1 = iota //此时，iota为2
	)
	fmt.Println(r1, r2, r3)

	const (
		t1, t2 = iota + 1, iota + 2 //t1:1, t2:2
		t3, t4 = iota + 3, iota + 4 //t3:4, t4:5
	)
	fmt.Println(t1, t2, t3, t4)
}
