package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 888
}

// 返回多个返回值，有形参名称
func foo3(a string, b int) (x int, y int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的变量赋值
	x = 1000
	y = 999
	return
}

func foo4(a string, b int) (x, y int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的变量赋值
	x = 888
	y = 666
	return
}

func main() {
	c := foo1("abc", 55)
	fmt.Println("c = ", c)

	x, y := foo2("abc", 66)
	fmt.Println("x = ", x, ", y = ", y)

	x1, y1 := foo3("abc", 66)
	fmt.Println("x1 = ", x1, ", y1 = ", y1)

	x2, y2 := foo4("abc", 66)
	fmt.Println("x2 = ", x2, ", y2 = ", y2)

}
