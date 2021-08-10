package main

import "fmt"

// const 枚举
const (
	BEIJING   = 1
	SHANGHAI  = 2
	SHENZHEN  = 3
	GUANGZHOU = 4
)

const (
	// 可以在 const 中添加一个关键字 iota，每行的iota会自动加1，第一行的iota默认值是0
	MONDAY = 10 * iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
)

const (
	// iota 只能配合 const 修饰符一起使用
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)

	// 编译不通过，常量是不允许修改的
	// length = 100

	fmt.Println(BEIJING)
	fmt.Println(FRIDAY)
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
