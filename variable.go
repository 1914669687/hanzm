package main

import "fmt"

//变量声明后必须使用

func main() {
	var s1 string
	var (
		name string
		age  int
		isOk bool
	)
	s1 = "variable"
	name = "hanzm"
	age = 22
	isOk = true
	fmt.Println(s1, name, age, isOk)

	var major = "计算机科学"
	tel := "123456" //简短声明变量只能在函数里使用
	fmt.Println(major, tel)

	var wechat, _ = "W123456", 1914 //在使用多重赋值时，匿名变量用一个"_"表示，不分配命名空间，也不会分配内存
	fmt.Println(wechat)
}
