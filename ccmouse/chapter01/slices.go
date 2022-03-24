package main

import "fmt"

// slice 是 array 的一个视图
// slice 中的结构有 ptr, len, cap，所以 slice 是可以向后扩展的
// 添加元素如果超过 cap，那么 slice 底层会重新分配一个更大的 array

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	// s := arr[2:6]
	// fmt.Println(s)
	// updateSlice(s)
	// fmt.Println("After UpdateSlice")
	// fmt.Println(s)

	s1 := arr[2:6]
	fmt.Println("s1=", s1)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := arr[3:5]
	fmt.Println("s2=", s2)

}
