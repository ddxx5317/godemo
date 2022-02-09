package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	n,err := strconv.Atoi("123")
	if err != nil{
		fmt.Println("error",err)
	}else {
		fmt.Println("result:",n)
	}

	n2 := strconv.FormatInt(16,2)
	fmt.Println("n2:",n2)

	b := strings.Contains("abc","a")
	fmt.Println("b:",b)

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Day())

	//时间日期
	ymdHms := now.Format("2006-01-02 15:04:05")
	ymd := now.Format("2006-01-02")
	hms := now.Format("15:04:05")
	year := now.Format("2006")
	fmt.Println(ymdHms)
	fmt.Println(ymd)
	fmt.Println(hms)
	fmt.Println(year)


}
