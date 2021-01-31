/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/31 09:53
 */
package main

import "fmt"

//if条件判断
func main() {
	//	age := 19
	//	if age > 18 { //如果age>18就执行这个{}中的代码
	//		fmt.Println("澳门首家线上赌场开业啦！")
	//	} else { //否者就执行这个{}中的代码
	//		fmt.Println("改写寒假作业啦！")
	//	}

	//多个判断条件
	//if age > 35 {
	//	fmt.Println("人到中年，多喝枸杞")
	//} else if age > 18 {
	//	fmt.Println("青年")
	//} else {
	//	fmt.Println("好好学习")
	//}

	//作用域，在if语句中的声明变量，只在if语句之内起作用，{ }之外是不起作用的
	if age := 19; age <= 18 {
		//否者就执行这个{}中的代码
		fmt.Println("改写寒假作业啦！")
	} else {
		//如果age>18就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业啦！")
	}
	//fmt.Println(age) //所以这里如果直接打印age这个变量，会直接提示undefined:age找不到这个变量声明
}
