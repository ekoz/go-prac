package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func whileTrue() {
	// go 语言的 while 死循环
	for {

	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		// 13除以2，被6整除，余1，首位写1
		// 6 除以2，被3整除，余0，第二位写0
		// 3 除以2，被1整除，余1，第三位写1
		// 1 除以2，被0整除，余1，第四位写1
		// 最后把 1011 反转成 1101，就是13的二进制
		convertToBin(13), // 1011 -> 1101
		convertToBin(100),
		// 0 是空字符串
		convertToBin(0),
	)
}
