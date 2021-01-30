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
	//for _, c := range s { //从字符串中拿出具体的字符
	//	fmt.Printf("%c\n", c)

	//"hello" ==> 'h' 'e' 'l' 'l' 'o'
	//字符串修改
	s2 := "白萝卜"             //分解成 '白'  '萝'  '卜'
	s3 := []rune(s2)        //把字符串强制转换成一个rune切片
	s3[0] = '红'             //双引号表示字符串，单引号表示字符
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	//双引号和单引号表示类型可不一样
	c1 := "红" //双引号包括就是字符串类型
	c2 := '红' //一个中文字占3个字节，一个字节占8位，3*8=24位，int32，是因为UTF-8前缀表示我是rune类型，所以24+8=32，rune(int32别名就是rune）
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4) //打印10进制的数值
}
