package main

// 注意，这种引入包的方式，只能用于 go env -w GO111MODULE=off 的方式
// 感谢b站上的字幕黨
// 需要了解 gomod
import (
	// _ 是别名的方法导入包，如果只希望执行包中的init方法，就可以将别名设置为 _，此时不会报错
	_ "ibothub.com/go-prac/aceld/chapter02/lib3"
	// mylib2 是别名
	"ibothub.com/go-prac/aceld/chapter02/lib1"
	mylib2 "ibothub.com/go-prac/aceld/chapter02/lib2"
	// 采用 . 的方式导入包，相当于将该包中所有的public方法都导入，
	// 不推荐使用，ide会提示 should not use dot imports (ST1001)
	. "ibothub.com/go-prac/aceld/chapter02/lib4"
)

func main() {
	lib1.GetLib1()
	mylib2.GetLib2()
	// 注意 lib4 的调用方式
	GetLib4()
}
