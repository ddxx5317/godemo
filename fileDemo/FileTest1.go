package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "D://test.txt"
	file,err :=os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("err=%v",err)
		return
	}

	str := "hello,world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	defer file.Close()
}
