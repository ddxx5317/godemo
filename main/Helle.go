package main

import "fmt"

func main() {
	fmt.Println("hello world!!!")
	var a int = 3
	var b int = 4
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=",a,"b=",b)

	var name string
	var age int

	fmt.Println("input a num:")
	fmt.Scanln(&age)
	fmt.Println("input num is :" ,age)


	fmt.Println("input a name and age, with space split")
	fmt.Scanf("%s %d",&name,&age)
	fmt.Println("name:",name,"age:",age)
}
