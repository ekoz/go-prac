package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "abc"

// 函数体外部必须使用 var 声明，外面声明的变量不是一个全局变量，而是一个包范围的变量
// dd := 5

// 简略写法
var (
	ii = 3
	jj = "abc"
	kk = true
)

func variableZeroValue() {
	var a int
	var s string

	// %q Quotation 的缩写，可以把引号打印出来，如果这里用的是 %s，那么无法打印出来
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	// int 和 string 可以不用写，golang会推断出类型
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableInitValueDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func euler() {
	// 不能用 4*i，这样 golang会认为i是一个变量，写成 4i，golang就能识别成复数
	c := 3 + 4i
	// cmplx.Abs(c): 5
	// cmplx.Pow(math.E, 1i*math.Pi): (0+1.2246467991473515e-16i)
	// cmplx.Exp(1i * math.Pi): (0+1.2246467991473515e-16i)
	fmt.Printf("cmplx.Abs(c): %v\n", cmplx.Abs(c))

	fmt.Printf("cmplx.Pow(math.E, 1i*math.Pi): %v\n", cmplx.Pow(math.E, 1i*math.Pi)+1)

	fmt.Printf("cmplx.Exp(1i * math.Pi): %v\n", cmplx.Exp(1i*math.Pi)+1)

	// cmplx.Pow(math.E, 1i*math.Pi): (0.000+0.000i)
	// cmplx.Exp(1i * math.Pi): (0.000+0.000i)
	fmt.Printf("cmplx.Pow(math.E, 1i*math.Pi): %0.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)

	fmt.Printf("cmplx.Exp(1i * math.Pi): %0.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// golang 必须显示转int
	// float64 浮点数是不准的，可能会得到 4.99999999999999
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("c: %v\n", c)
}

func consts() {
	// const filename string = "abc.txt"
	// const a, b = 3, 4

	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// 普通枚举类型
	// const (
	// 	cpp    = 0
	// 	java   = 1
	// 	python = 2
	// 	golang = 3
	// )
	const (
		// 自增值枚举类型
		cpp = iota
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// 1 1024 1048576 1073741824 1099511627776 1125899906842624
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitValue()
	variableInitValueDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
}
