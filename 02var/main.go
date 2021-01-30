/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 12:53
 */
package main

import "fmt"

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go语言中变量声明必须使用，不使用就编译不过去
	fmt.Print(isOk)             //在终端中输出要打印的内容
	fmt.Printf("name：%s", name) //在终端输出打印的内容
	fmt.Println(age)            //打印完指定的内容之后，在后面加一个换行符
	s3 := "哈哈哈"
	fmt.Println(s3)
}
