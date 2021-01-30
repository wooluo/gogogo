/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 18:42
 */
package main

//字符串 \\本来是具有特殊含义的，我应该告诉程序我写的\就是个单纯的\

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"/Users/zl/go/src/github.com/wooluo/gogogo/08string\""
	fmt.Println(path)

	s := "I'm Ok"
	fmt.Println(s)

	//多行的字符串
	s2 := `
	世情薄
	人情恶
	雨送黄昏花易落`
	fmt.Println(s2)
	s3 := `//Users//zl//go/src//github.com//wooluo//gogogo//08string`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	//fmt.Printf("%s%s", name, world)
	fmt.Println(ss1)

	//分隔
	ret := strings.Split(s3, "/")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
}
