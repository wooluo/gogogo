/*
*File  : main.go
*Author: wooluo
*Date  : 2021/2/2 15:51
 */
package main

import (
	"fmt"
)

var num int64 = 10

//func testGlobalVar() {
//	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
//}
func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num)
}
func main() {
	testNum()
}

func testLocalVar2(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100 //函数的参数也是只在本函数中生效
		fmt.Println(z)
	}
	//fmt.Println(z)
	//
}
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}
