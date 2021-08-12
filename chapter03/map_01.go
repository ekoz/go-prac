package main

import "fmt"

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap 是一个空map")
	}

	myMap = make(map[string]string, 3)
	myMap["one"] = "Java"
	myMap["two"] = "C++"
	myMap["three"] = "Python"
	fmt.Println(myMap)

	myMap1 := make(map[int]string)
	myMap1[0] = "Java"
	myMap1[1] = "C++"
	myMap1[2] = "Python"
	fmt.Println(myMap1)

	myMap2 := map[int]string{
		0: "php",
		1: "golang",
		2: "erlang",
	}
	fmt.Println(myMap2)

}
