/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/31 10:31
 */
package main

import "fmt"

//for循环
func main() {
	//基本格式
	//for i := 0; i < 1; i++ {
	//	fmt.Println(i)
	//}

	//变种1
	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println("i")
	//}

	//变种2
	//var i = 5
	//for i < 10 {
	//	fmt.Println(i)
	//	i++ //从5开始累加，直到小于10的整数
	//}

	//for {
	//	fmt.Println("123") //听说你电脑很不错， 可以跑一下试试
	//}

	//for range循环
	s := "1"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
