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


}
