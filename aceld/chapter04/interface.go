package main

import "fmt"

// interface{} 是万能数据类型
// int string float32 float64 struct 都实现了 interface{}
// 可以用 interface{} 引用任意的数据类型

func myFunc(args interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(args)

	// interface{} 到底是如何区分此时底层的引用类型到底是什么呢？

	// 给 interface{} 提供类型断言机制

	value, ok := args.(string)
	if ok {
		fmt.Println("args is string, value is", value)
	} else {
		fmt.Println("args is not string")
	}
}

// 首字母小写，只能当前类中调用
type book struct {
	name string
}

func main() {
	myBook := book{"Java编程思想"}
	myFunc(myBook)

	s := "abc"
	myFunc(s)

	i := 10
	myFunc(i)
}
