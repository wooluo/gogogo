/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 21:27
 */
package main

import (
	"fmt"
)

//fmt占位符
func main() {
	var n = 100 //定义变量
	//查看类型
	fmt.Printf("%T\n", n) //查看字符类型
	fmt.Printf("%v\n", n) //查看数值
	fmt.Printf("%b\n", n) //查看二进制
	fmt.Printf("%d\n", n) //查看十进制
	fmt.Printf("%o\n", n) //查看八进制
	fmt.Printf("%x\n", n) //查看字符串
	var s = "Hello 海淀！"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s) //增加#,打印的对象是什么类型的，这层输出是字符串就会增加个双引号

}
