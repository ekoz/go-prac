package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close 可以关闭一个channel
		// 关闭 channel 后，可以继续
		close(c)
	}()

	for {
		// ok 如果为true表示channel没有关闭，如果为false标识已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("主进程执行完毕")
}
