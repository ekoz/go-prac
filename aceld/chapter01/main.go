package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Printf("你好，世界")
	fmt.Println(int64(time.Now().Unix()))
	fmt.Println(int64(time.Now().Unix()))
	fmt.Println(int64(time.Now().Unix()))
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Nanosecond())
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	result1, _ := rand.Int(rand.Reader, big.NewInt(100))
	result2, _ := rand.Int(rand.Reader, big.NewInt(100))
	result3, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(result, result1, result2, result3)
}
