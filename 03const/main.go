/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 14:19
 */
package main

import "fmt"

//const pi = 3.1415926
//
////批量声明常量
//const (
//	statusOK = 200
//	notFound = 404
//)
//const (
//	n1 = 100
//	n2
//	n3
//)

//iota
//const (
//	a1 = iota // 0
//	a2        //1
//	a3        //2
//)
//const (
//	b1 = iota // 0
//	b2 = iota //1
//	_  = iota
//	b3 = iota //2
//)
//
//const (
//	c1 = iota //0
//	c2 = 100  //100
//	c3 = iota
//	c4 = iota
//)
const ( //定义数量级
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
}
