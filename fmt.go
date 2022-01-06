package main

import (
	"fmt"
)

func main() {
	var n = 100
	fmt.Printf("%d\n", n)
	fmt.Printf("二进制: %b\n", n)
	fmt.Printf("十六进制: %x\n", n)
	fmt.Printf("value: %v\n", n)
	fmt.Printf("类型: %T\n", n)
	var s = "hanzm"
	fmt.Printf("%s\n", s)
	fmt.Printf("%#v\n", s) //如果是字符串就会加上引号
	var s1 = "I love you"
	var s2 = fmt.Sprintf("%s %s", s, s1)
	fmt.Println(s2)
}
