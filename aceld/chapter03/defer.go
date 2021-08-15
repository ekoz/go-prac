package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called..")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called..")
	return 0
}

// return 比 defer 要先输出
// return func called..
// defer func called..
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer 类似于 java 中的final，当前函数生命周期结束之后才会出栈，函数体退出前执行，类似于压栈，先入后出
	defer fmt.Println("main end")
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	returnAndDefer()
}
