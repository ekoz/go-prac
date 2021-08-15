package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float32  `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie1 := movie{"戏剧之王", 1999, 30.0, []string{"周星驰", "张柏芝", "莫文蔚"}}
	fmt.Printf("movie is %v\n", movie1)

	movieStr, err := json.Marshal(movie1)
	if err != nil {
		fmt.Printf("json format error, reason: %v", err)
		return
	}
	fmt.Printf("json string is %s\n", movieStr)

	// 将 json string 转成 结构体
	movie2 := movie{}
	err = json.Unmarshal(movieStr, &movie2)
	if err != nil {
		fmt.Printf("json unmarshal error, reason: %v", err)
		return
	}
	fmt.Printf("movie2 is %v\n", movie2)

}
