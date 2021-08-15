package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine exit

func main() {

	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			// go 退出
			runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Printf("a + b = %d\n", (a + b))
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
