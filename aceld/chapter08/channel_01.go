package main

import (
	"fmt"
)

// channel 分为有缓冲和无缓冲
// 当前用例是无缓冲，主线程一定会等待子线程的信号发送完毕，接收到信号后，才会继续执行
// 有缓冲channel 可以用 epoll 来理解，每个 channel 类似于队列

func main() {
	c := make(chan int)

	// go func() {
	// 	defer fmt.Println("子线程.Defer")
	// 	fmt.Println("子线程执行完毕..")
	// 	c <- 666
	// }()

	// <-c 必须连在一起，不能分开
	// num := <-c
	num := 0
	fmt.Printf("c: %v\n", len(c))

	fmt.Printf("主线程接收到子线程信号\"%d\"，主线程执行完毕", num)
}
