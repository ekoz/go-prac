package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int // zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	s1 := make([]int, 16)
	printSlice(s1)
	s2 := make([]int, 10, 32)
	printSlice(s2)

	fmt.Println("Copy Slice")
	copy(s2, s)
	printSlice(s2)

	fmt.Println("Delete Slice Element")
	s2 = append(s2[:3], s1[3:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
