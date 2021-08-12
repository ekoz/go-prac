package lib1

import "fmt"

// 方法名大写 - public，小写 private
// 对外提供就采用首字母大写
func GetLib1() {
	fmt.Println("lib1 return()")
}

func init() {
	fmt.Println("lib1 init success...")
}
