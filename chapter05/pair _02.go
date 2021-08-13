package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path, _ := os.Getwd()
	fmt.Printf("path: %v\n", path)

	// file: pair<type:*os.File, value:"../docs/hello.txt"文件描述符>
	file, err := os.OpenFile(path+"/docs/hello.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error,", err)
		return
	}

	// r: pair<type:, value:>
	var r io.Reader
	// r: pair<type:*os.File, value:"../docs/hello.txt"文件描述符>
	r = file

	// w: pair<type:, value:>
	var w io.Writer
	// w: pair<type:*os.File, value:"../docs/hello.txt"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("Hello, Hero!\n"))
}
