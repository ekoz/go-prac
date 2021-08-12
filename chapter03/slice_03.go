package main

import "fmt"

// b站字幕上，有网友提示说当容量超出时，实际情况是重新申请一块内存，然后拷贝，未经证实，仅供参考
// 并非总是扩容2倍，对于长度比较大的切片，会尝试扩容1/4，已节省空间

func printArray(arr []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(arr), cap(arr), arr)
}

func main() {
	// 3 标识数组长度，5标识数组容量，当 append 数组append长度超过5时，则自动扩容，每次扩容都是5的倍数
	var numbers = make([]int, 3, 5)
	printArray(numbers)

	// 向 numbers 追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	// 向 numbers 追加一个元素2
	numbers = append(numbers, 2)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	// 向 numbers 追加一个元素3，此时，由于len超过cap，所以cap开始扩容
	numbers = append(numbers, 3)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	// numbers2 的 len 和 cap 都是3，当长度超过3时，自动扩容至6，依此类推
	var numbers2 = make([]int, 3)
	printArray(numbers2)

	// 开始扩容
	numbers2 = append(numbers2, 1)
	printArray(numbers2)

	arr := []int{1, 2, 3, 4, 5}

	arr1 := arr[1:]
	arr2 := arr[0:2]

	// 此时 arr 和 arr2 的第一个元素都会修改为100
	arr[0] = 100
	fmt.Println(arr, arr1, arr2)

	arr3 := make([]int, len(arr))
	copy(arr3, arr)
	// 此时，只会修改arr第一个元素，arr3 的第一个元素不受影响
	arr[0] = 10
	fmt.Println(arr, arr3)

}
