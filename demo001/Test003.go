package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) run(name string) string{
	fmt.Println(name + " run......")
	return "MKing Good"
}
func main() {

	//方式1
	var person Person
	person.Name="King"
	person.Age=20
	fmt.Println(person)

	//方式2
	person2 := Person{"King",20}
	fmt.Println(person2)


	//方式3
	var person3 *Person = new(Person)
	//(*person3).Name="DX"
	//(*person3).Age=41
	//简化写法
	person3.Name="king2"
	person3.Age=41
	fmt.Println(*person3)


	//方式4
	var person4 *Person = &Person{"king99",20}
	fmt.Println(*person4)


	var personTag Person
	personTag.Name="KingTag"
	personTag.Age=200
	fmt.Println(personTag)
	jsonStr, err := json.Marshal(personTag)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("json:",string(jsonStr))


	//方法调用
	name := personTag.run("king")
	fmt.Println("name:",name)


	var person999 = &Person{"KO",30}
	fmt.Println(*person999)

}
