package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string `info:"name" doc:"我的名字"`
	Age  int    `info:"age"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info: %s, doc: %s\n", info, doc)
	}
}

func main() {
	user := &user{"小明", 10}
	findTag(user)
}
