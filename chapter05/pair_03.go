package main

import (
	"fmt"
	// "os"
)

// 接口
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
	name string
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book named", this.name)
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book named", this.name)
}

func main() {
	// book: pair<type:Book, value:book{}地址>
	book := &Book{"Java编程思想"}

	// r: pair<type:, value:>
	var r Reader
	// r: pair<type:Book, value:book{}地址>
	r = book
	r.ReadBook()

	// w: pair<type:, value:>
	var w Writer
	// 此处为什么能断言成功？因为 w r 具体的type是一致的
	// w: pair<type:Book, value:book{}地址>
	w = r.(Writer)
	w.WriteBook()

}
