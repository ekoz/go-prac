package main

import "fmt"

type Book struct {
	title  string
	auther string
	price  float32
}

func changeBook(book Book) {
	book.auther = "匿名"
}

func changeBook2(book *Book) {
	book.auther = "匿名2"
}

func main() {
	var book1 Book
	book1.title = "Golang 从入门到放弃"
	book1.auther = "佚名"
	book1.price = 8.8

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
