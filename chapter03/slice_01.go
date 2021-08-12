package main

import "fmt"

func printArray(arr []int) {
	// 引用传递
	for _, value := range arr {
		fmt.Println("value is ", value)
	}

	arr[0] = 100
}

func main() {
	// 动态数组，切片 slice
	myarr := []int{1, 2, 3}
	fmt.Printf("myarr type is %T\n", myarr)
	printArray(myarr)

	fmt.Println("========")

	for _, value := range myarr {
		fmt.Println("value is ", value)
	}
}
