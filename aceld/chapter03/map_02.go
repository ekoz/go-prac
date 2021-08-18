package main

import "fmt"

func main() {
	myMap := make(map[int]string)
	myMap[0] = "BeiJing"
	myMap[1] = "ShangHai"
	myMap[2] = "GuangZhou"
	myMap[3] = "ShenZhen"

	// 遍历
	for index, item := range myMap {
		fmt.Println("index is", index, ", value is", item)
	}

	// 删除
	delete(myMap, 1)

	// 修改
	myMap[3] = "WuHan"

	fmt.Println("======================")
	// 遍历
	for index, item := range myMap {
		fmt.Println("index is", index, ", value is", item)
	}

	fmt.Println(myMap[5] == "", myMap[10] == "", myMap[0] == "BeiJing")
}
