package main

// 当channel已经满了，再向里面写数据，就会阻塞
// 当channel为空，从里面取数据也会阻塞

import (
	"fmt"
	"time"
	// "runtime"
)

func main() {
	// 带有三个元素缓冲的channel
	chanLen := 3
	c := make(chan int, chanLen)

	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子进程.Defer")

		for i := 0; i < chanLen+1; i++ {
			c <- i
			fmt.Printf("子进程正在运行，len(c)=%d, cap(c)=%d\n", len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < chanLen+1; i++ {
		num := <-c
		fmt.Println("num is", num)
	}

	fmt.Println("主进程执行完毕")
}
