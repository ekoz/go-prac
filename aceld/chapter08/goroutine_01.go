package main

import (
	"fmt"
	"time"
)

func ibotTask() {
	i := 0
	for {
		i++
		fmt.Printf("ibot Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个go程，去执行 ibotTask
	go ibotTask()

	// 主线程死循环
	for {
	}
}
