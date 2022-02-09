package main

import (
	"fmt"
)

func main() {
	//数组
	arr := [5]int {11, 22, 33, 44, 55}
	for i := 0; i < len(arr); i++{
		fmt.Printf("arr[%d]=%d\n", i, arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	//切片使用方法1
	slice1 := arr[1:3]
	fmt.Println("sl:",slice1)
	fmt.Println("sl:",cap(slice1))

	//切片使用方法2
	var slice2 []int = make([]int,4)
	slice2[0] = 22
	slice2[2] = 55
	fmt.Println(slice2)
	slice2 = append(slice2,88)

	slice2 = append(slice2,slice2...)
	fmt.Println(slice2)


	str := "abc@com"
	fmt.Println(str[4:])

	//映射
	//var dict map[string]string //定义dict为map类型
	//dict = make(map[string]string) //让dict可编辑（分配内存空间）
	//
	//dict["a"] = "苹果"
	//dict["b"] = "香蕉"



	dict := map[string]string{"a":"苹果","b":"香蕉"}

	value, ok := dict["a"] //ok是看当前key是否存在返回布尔，value返回对应key的值
	if ok {
		fmt.Println("存在",value)
	}else{
		fmt.Println("不存在")
	}
	for k,v := range dict{
		fmt.Println(k,"=",v)
	}

	//map切片

}
