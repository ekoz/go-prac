package main

import (
	"fmt"
	"sync"
	"time"
)

type wgService struct {
	wg sync.WaitGroup
}

var wgs wgService
var ch chan int

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // Done()用来表示goroutine已经完成了，减少一次计数器
}

func process2(i int) {
	defer wgs.wg.Done() // Done()用来表示goroutine已经完成了，减少一次计数器
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	fmt.Printf("ch: %v, len(ch): %v\n", ch, len(ch))
	<-ch
}

func main() {
	ch = make(chan int, 2)
	no := 10
	for i := 0; i < no; i++ {
		wgs.wg.Add(1)
		ch <- 1
		go process2(i)
	}
	wgs.wg.Wait() // Wait()用来等待所有需要等待的goroutine完成。
	fmt.Println("All go routines finished executing")
}
