package main

import "fmt"

// map 使用hash表，必须可以比较相等
// 除了 slice，map，function的内建类型都可以作为key
// Struct 类型不包含上述字段，也可做为key

func main() {

	// map 是无序的
	m := map[string]string{
		"name":   "ekozhan",
		"age":    "35",
		"sex":    "man",
		"course": "golang",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	courseName := m["course"]
	fmt.Println(courseName)

	key := "course1"
	if courseName, ok := m[key]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Printf("key [%s] is not exist\n", key)
	}

	fmt.Println("====== delete ======")
	name, ok1 := m["name"]
	fmt.Println(name, ok1)

	delete(m, "name")
	name, ok1 = m["name"]
	fmt.Println(name, ok1)
}
