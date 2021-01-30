/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 21:44
 */
package main

import (
	"fmt"
)

//布尔值
func main() {
	b1 := true
	var b2 bool                         //默认是flase
	fmt.Printf("%T\n", b1)              //T是指类型
	fmt.Printf("%T value:%v\n", b1, b2) //v可以指任意变量的值
}
