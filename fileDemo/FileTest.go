package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file,err := os.Open("D://password.txt")
	if err != nil{
		fmt.Println("file open error",err)
	}
	fmt.Printf("file=%v\n",file)

	reader := bufio.NewReader(file)
	for  {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("over")

	defer file.Close()
}
