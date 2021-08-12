package lib2

import "fmt"

// 方法名大写 - public，小写 private
// 对外提供就采用首字母大写
func GetLib2() {
	fmt.Println("lib2 return()")
}

func init() {
	fmt.Println("lib2 init success...")
}
