package main

import "fmt"

// 类名首字母大写，其他包也可以引入
type Hero struct {
	Name  string
	Age   int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	hero := Hero{Name: "美国队长", Age: 100, Level: 0}
	fmt.Printf("%v\n", hero)

	hero.SetName("蜘蛛侠")
	fmt.Printf("%v\n", hero)
}
