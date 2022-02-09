package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int,3)
	fmt.Println(intChan)
	fmt.Println(&intChan)

	intChan <- 10
	num := 211
	intChan <- num
	fmt.Println(intChan)

	num2 := <- intChan
	fmt.Println(num2)
	num3 := <- intChan
	fmt.Println(num3)

	close(intChan)
}
