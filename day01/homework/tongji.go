/*
*File  : tongji.go
*Author: wooluo
*Date  : 2021/1/30 22:44
 */
package main

import "fmt"

//作业内容，统计s1字符串中的英文有多少个
//01func main() {
//	a := 0
//	s1 := "hhhhhello沙河小王子"
//	for _, i := range s1 {
//		if i < 'z' { //包括字母内
//			a++
//		}
//	}
//	fmt.Println(a) //等于9，说明字符串中英文字母有9个
//}

//作业内容，统计s1字符串中的汉字有多少个
func main() {
	a := 0
	s1 := "hhhhhello沙河小王子王子王子王子王子王子王子"
	for _, i := range s1 {
		if i > 'z' { //大于字母
			a++
		}
	}
	fmt.Println(a) //等于17，说明有17个汉字
}
