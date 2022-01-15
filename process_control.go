package main

import (
	"fmt"
)

func main() {
	a := 100
	//a++, a--单独成一句话，不能写在其他表达式中，且++a, --a都不合法
	a++
	fmt.Println(a)

	b := true
	c := false
	fmt.Println(b || c, b && c)

	b1 := 4                          //0100
	c1 := 8                          //1000
	fmt.Println(b1&c1, b1|c1, b1^c1) //0000, 1100, 1100(相反为1，相同为0)

	//golang中的if条件必须为bool类型
	//大括号必须存在，即使只有一条语句，左括号必须与if else在同一行
	//if else golang的特殊用法
	if age := 22; age > 18 { //age的作用域仅限于if else
		fmt.Printf("您%d岁, 已成年\n", age)
	} else {
		fmt.Printf("您%d岁, 未成年\n", age)
	}

	//switch语句
	day := 6
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	}
	//case可以是表达式
	switch {
	case day > 3:
		fmt.Println("一周过去一半了")
	default:
		fmt.Println("本周才刚开始")
	}
	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("111")
	case 6:
		fmt.Println("333")
		fallthrough
	case 7:
		fmt.Println("123")
	default:
		fmt.Println("error")
	}

	//golang只有for循环，没有while do-while
	for i := 1; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
	var arr = []int{1, 2, 3, 4}
	for key, val := range arr { //for range相当于python中的enumerate
		fmt.Println(key, val)
	}
}
