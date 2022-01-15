package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var s1 = `
	春水初生
	春林初盛
` //多行字符串用 ``
	s2 := "春风十里不如你"

	//字符串拼接
	s3 := s1 + s2
	fmt.Println(s3)
	s3 = fmt.Sprintf("%s\t%s", s1, s2)
	fmt.Printf("fmt.Sprintf: %s\n", s3)
	s3 = strings.Join([]string{s1, s2}, "\t")
	fmt.Printf("string.Join: %s\n", s3)
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString("\t")
	buffer.WriteString(s2)
	fmt.Printf("bytes.Buffer: %v\n", buffer.String())

	//字符串切片
	s4 := "hello world!"
	fmt.Printf("%c\n", s4[0])
	fmt.Println(s4[:5])
	fmt.Println(s4[6:])
	fmt.Println(s4[2:4])

	//其他操作
	fmt.Printf("length: %v\n", len(s4))
	fmt.Printf("strings.Split: %v\n", strings.Split(s4, " "))
	fmt.Println(strings.Contains(s4, "!"))
	fmt.Println(strings.HasPrefix(s4, "h"))
	fmt.Println(strings.HasSuffix(s4, "d"))

}
