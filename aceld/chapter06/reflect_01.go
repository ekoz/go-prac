package main

import (
	"fmt"
	"reflect"
)

func reflectNum(args interface{}) {
	fmt.Println("type:", reflect.TypeOf(args), ", value:", reflect.ValueOf(args))
}

func main() {
	// ValueOf TypeOf
	var num float64 = 3.1415926
	reflectNum(num)
}
