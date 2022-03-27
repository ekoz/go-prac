package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// UTF-8
	// 中文占用3字节
	s := "hello,你好世界!"
	fmt.Println(len(s))

	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Printf("rune count is %d\n", utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
