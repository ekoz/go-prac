package main

import "fmt"

// 继承

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat(food string) {
	fmt.Println("Human eat", food)
}

func (this *Human) Walk() {
	fmt.Println("Human walk..")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Walk() {
	fmt.Println("SuperMan walk..")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan fly..")
}

func main() {
	human := Human{"张三", "男人"}
	human.Eat("香蕉")
	human.Walk()

	superMan := SuperMan{Human{"李四", "超人"}, 0}
	superMan.Fly()
	fmt.Printf("%v\n", superMan)
}
