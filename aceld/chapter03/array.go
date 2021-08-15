package main

import "fmt"

func main() {
	var myarr [10]int

	myarr1 := [10]int{1, 2, 3, 4}
	myarr2 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myarr); i++ {
		fmt.Println(i, myarr[i])
	}

	for index, value := range myarr1 {
		fmt.Println("index =", index, "value =", value)
	}

	fmt.Printf("myarr type is %T\n", myarr)
	fmt.Printf("myarr1 type is %T\n", myarr1)
	fmt.Printf("myarr2 type is %T\n", myarr2)
}
