package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"modtest/protobuf/modtest/protobuf"
)

func main()  {
	//序列化
	d := protobuf.User{
		Name: "jahon",
		Male: true,
	}
	res, err := proto.Marshal(&d)
	fmt.Println("res", res, err)
	//反序列化
	obj := &protobuf.User{}
	err = proto.Unmarshal(res, obj)
	if err != nil {
		panic(err)
	}
	fmt.Println("obj", obj, err)
}
