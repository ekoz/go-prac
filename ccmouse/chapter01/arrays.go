package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	// 将100赋值给第0个元素，因为array是值传递，所以内部 arr[0] 能赋值，但是外面访问的时候，arr[0] 还是老数据
	arr[0] = 100
}

func printArray2(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func main() {
	// 数组是值类型
	var arr1 [5]int

	arr2 := [3]int{1, 3, 5}

	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int

	fmt.Println(grid)

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[i])
	// }

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	printArray(arr1)
	// 编译不通过，因为方法参数需要一个长度为5的int数组
	// printArray(arr2)
	printArray(arr3)
	printArray2(&arr3)
}
