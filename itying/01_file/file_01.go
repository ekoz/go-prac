package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 判断文件是否存在
	_, err := os.Stat("./Monday1.txt")
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("file stat error: %v\n", err)
		fmt.Println("文件 ./Monday1.txt 不存在")
	}

	// 创建文件
	f, err := os.Create("./Monday.txt")
	if err != nil {
		fmt.Printf("file create error: %v\n", err)
		return
	}
	absPath, err := filepath.Abs(f.Name())
	if err != nil {
		fmt.Printf("file abs path error: %v\n", err)
		return
	}
	dirPath := filepath.Dir(f.Name())
	fmt.Printf("文件创建成功\n绝对路径：%s\n相对路径：%s\n", absPath, dirPath)

	// 创建目录
	// 可重复执行，如果已经创建，则不会重复创建
	err = os.MkdirAll("./a1/a2/a3", 0755)
	if err != nil {
		fmt.Printf("mkdir all error: %v\n", err)
		return
	}

	// 创建目录-2
	// 不可重复执行，如果已经存在该目录，则报错
	err = os.Mkdir("b1", 0755)
	if err != nil {
		fmt.Printf("mkdir error: %v\n", err)
	}

	os.RemoveAll("a1")
	os.RemoveAll("b1")
}
