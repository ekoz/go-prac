package main

import "fmt"

func printArray(arr []int) {
	// 引用传递
	fmt.Printf("param type is %T, length is %d\n", arr, len(arr))
	for _, value := range arr {
		fmt.Println("value is ", value)
	}

}

func main() {
	var slice0 []int

	//声明 slice 是一个切片，并且初始化，默认值是1，2，3，长度len是3
	slice1 := []int{1, 2, 3}
	printArray(slice1)

	fmt.Println("============")

	var slice2 []int
	printArray(slice2)

	// 给 slice2 开辟3个空间，默认值都是0
	slice2 = make([]int, 3)
	printArray(slice2)

	// 声明一个 slice3，并给 slice3 开辟3个空间，默认值都是0
	var slice3 []int = make([]int, 3)
	printArray(slice3)

	// 通过 := 推导 slice4 是一个切片
	slice4 := make([]int, 3)
	printArray(slice4)

	if slice0 == nil {
		fmt.Println("slice0 是一个空切片")
	} else {
		fmt.Println("slice0 是有空间的")
	}

}
