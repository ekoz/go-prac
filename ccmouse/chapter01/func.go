package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupport operation " + op)
	}
}

func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupport operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	// 快捷键可以自动返回 q r 变量名进行接收
	return a / b, a % b
}

func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	// 函数式编程
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(ns ...int) int {
	// 可变参数
	s := 0
	for i := range ns {
		s = s + ns[i]
	}
	return s
}

func swap(a int, b int) {
	// 值传递 和 引用传递
	// 值传递，拷贝传递
	// Golang 只有值传递
	a, b = b, a
}

func swap2(a *int, b *int) {
	*a, *b = *b, *a
}

func swap3(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Printf("eval(1, 3, \"+\"): %d\n", eval(1, 3, "+"))
	q, r := div2(4, 3)
	println(q, r)
	q2, _ := div3(8, 3)
	println(q2)

	if i, err := eval2(2, 3, "*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	fmt.Printf("%d\n", apply(pow, 2, 3))

	// 匿名函数
	fmt.Printf("%d\n", apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 3))

	fmt.Printf("sum(1, 2, 3, 4, 5): %d\n", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(a, b)
	println(a, b)
	// swap2(&a, &b)
	// println(a, b)
	a, b = swap3(a, b)
	println(a, b)
}
