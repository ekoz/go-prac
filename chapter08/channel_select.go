package main

import "fmt"

// change 和 range 配合使用
// 这个 sample 没读懂，哈哈，养肥了再看

func fibonacii(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果 c 可写，则该case会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacii(c, quit)

	fmt.Println("主进程执行完毕")
}
