package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "不及格"
	case score < 80:
		g = "中等"
	case score < 90:
		g = "良好"
	case score <= 100:
		g = "优秀"
	default:
		panic(fmt.Sprintf("err score %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	// open abc.txt: The system cannot find the file specified.
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(61),
		grade(83),
		grade(92),
		grade(100),
		// panic: err score 105

		// goroutine 1 [running]:
		// main.grade(0x669720)
		//         D:/GoWorks/ibothub.com/go-prac/ccmouse/chapter01/branch.go:20 +0xc8
		// main.main()
		//         D:/GoWorks/ibothub.com/go-prac/ccmouse/chapter01/branch.go:48 +0x173
		// exit status 2
		grade(105),
	)
}
