/*
*File  : main.go
*Author: wooluo
*Date  : 2021/1/30 21:55
 */
package main

import (
	"fmt"
)

//byte和rune类型
//Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型，这样的好处就是不会乱码了
func main() {
	s := "hello海淀해정."
	//len()求得是byte字节的数量
	n := len(s) //求字符串s的长度，把长度保存到变量n中
	fmt.Println(n)

	//	for i := 0; i <= len(s); i++
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n", s[i]) // %c:字符
	//}
	for _, c := range s { //从字符串中拿出具体的字符
		fmt.Printf("%c\n", c)
	}
}
