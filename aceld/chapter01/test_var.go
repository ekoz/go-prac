package main

import "fmt"

// 声明全局变量
var ga int
var gb int = 20
var gc = 30

// 全局变量声明无法使用 := 方法
// gd := 50

func main() {
	// 方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("type of bb = %T\n", bb)

	// 方法三：声明一个变量，省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 200
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("cc = ", cc)
	fmt.Printf("type of cc = %T\n", cc)

	//方法四：常用，省去var关键字，直接自动匹配
	d := 300
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	dd := "abcd"
	fmt.Println("dd = ", dd)
	fmt.Printf("type of dd = %T\n", dd)

	e := 3.14
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	fmt.Println("ga = ", ga)
	fmt.Printf("type of ga = %T\n", ga)
	fmt.Println("gb = ", gb)
	fmt.Printf("type of gb = %T\n", gb)
	fmt.Println("gc = ", gc)
	fmt.Printf("type of gc = %T\n", gc)
	// fmt.Println("gd = ", gd)
	// fmt.Printf("type of gd = %T\n", gd)

	// 声明多个变量
	var ma, mb int = 100, 200
	var mc, md = 100, "abcd"
	fmt.Println("ma = ", ma)
	fmt.Printf("type of ma = %T\n", ma)
	fmt.Println("mb = ", mb)
	fmt.Printf("type of mb = %T\n", mb)
	fmt.Println("mc = ", mc)
	fmt.Printf("type of mc = %T\n", mc)
	fmt.Println("md = ", md)
	fmt.Printf("type of md = %T\n", md)

	// 声明多个变量，方法二
	var (
		mx int  = 100
		my bool = false
	)
	fmt.Println("mx = ", mx)
	fmt.Printf("type of mx = %T\n", mx)
	fmt.Println("my = ", my)
	fmt.Printf("type of my = %T\n", my)
}
