package main

import "fmt"

func main() {
	var a string

	// pair<type:string, value:"aceld">
	a = "aceld"

	var allType interface{}

	allType = a

	str, _ := allType.(string)
	fmt.Printf("str: %v\n", str)
}
