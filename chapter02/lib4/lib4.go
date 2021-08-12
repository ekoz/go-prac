package lib4

import "fmt"

// 方法名大写 - public，小写 private
// 对外提供就采用首字母大写
func GetLib4() {
	fmt.Println("lib4 return()")
}

func init() {
	fmt.Println("lib4 init success...")
}
