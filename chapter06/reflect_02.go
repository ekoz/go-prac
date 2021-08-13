package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// this *User 注意，User 结构体的方法不是指针类型，如果是指针类型，NumMethod 的数量是 0
func (this User) ToString() {
	fmt.Printf("this: %v\n", this)
}

func DoFieldAndMethod(input interface{}) {
	// 获取 input 的 type
	t := reflect.TypeOf(input)
	// 获取 input 的 value
	v := reflect.ValueOf(input)
	fmt.Println("type is", t, ", type name is", t.Name(), ", value is", v)

	// 通过 type 获取里面的字段
	// 1、获取interface的reflect.Type，通过Type得到NumField，进行遍历
	// 2、得到每个field，数据类型
	// 3、通过field对应的interface()方法得到对应的value
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("field is %s, field type is %v, value is %v\n", field.Name, field.Type, value)
	}

	// 通过 type 获取里面方法
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("method: %v, method name is %s, method type is %v\n", method, method.Name, method.Type)
	}

}

func main() {
	user := User{1, "小明", 8}
	user.ToString()
	DoFieldAndMethod(user)
}
