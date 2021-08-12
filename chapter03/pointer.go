package main

import "fmt"

// 指针的作用，通俗来讲就是在函数内部修改外部的值

// 这是一个不带指针的普通方法
func swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
}

func runSwap() {
	var a, b int = 10, 20
	swap(a, b)
	fmt.Println("a =", a, "b =", b)
}

// 指针交换
func swapPointer(pa *int, pb *int) {
	tmp := *pa
	*pa = *pb
	*pb = tmp
}

func runSwapPointer() {
	var a, b int = 10, 20
	swapPointer(&a, &b)
	fmt.Println("a =", a, "b =", b)
}

func main() {
	runSwap()
	runSwapPointer()

	a := 10
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(&a)

	// 二级指针
	var pp **int
	pp = &p
	fmt.Println(pp)
	fmt.Println(&p)
}
