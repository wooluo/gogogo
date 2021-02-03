/*
*File  : 01func.go
*Author: wooluo
*Date  : 2021/2/2 15:35
 */
package main

import "fmt"

//01func intSum(x int, y int) int {
//	return x + y
//}

//01func intSum(x, y int) int { //可以简略类型，写成这样，因为这两个参数，类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型
//	return x + y
//}

func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
ret1 :intSum2()
ret1 :intSum2(10)
ret1 :intSum2(10,20)
ret1 :intSum2(10,20,30)
fmt.Println(ret1,ret2,ret3,ret4) //0 10 30 60


func sayHello() {
	fmt.Println("Hello 沙河")
}

func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)
}
