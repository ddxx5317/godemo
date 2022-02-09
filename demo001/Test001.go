package main

import (
	"fmt"
	"math/rand"
	"time"
)

func adder() func(int) int {
	var result = 10
	return func(i int) int {
		result = result + i
		return result
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(100) + 1
	fmt.Println(result)

	f := adder()
	fmt.Println("sum=",f(1))
	fmt.Println("sum=",f(2))
}
