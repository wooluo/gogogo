/*
*File  : chengfabiao.go
*Author: wooluo
*Date  : 2021/1/31 12:38
 */
package main

import (
	"fmt"
)

var b int //这里全局声明，如果放到程序入口里面会报错，提示没有使用，当然也可以不写声明，😝
func main() {
	for a := 1; a <= 9; a++ {
		for b := 1; b <= a; b++ { //如果写成b<=10的话，那么就把所有a*b全部打印出来，就不是斜梯式的图标了
			fmt.Printf("%d * %d =%d\t", b, a, b*a)
		}
		fmt.Println()
	}
}
