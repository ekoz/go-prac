package lib3

import "fmt"

// 方法名大写 - public，小写 private
// 对外提供就采用首字母大写
func GetLib3() {
	fmt.Println("lib3 return()")
}

func init() {
	fmt.Println("lib3 init success...")
}
