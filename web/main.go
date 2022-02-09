package main

import(
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎使用Go语言")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("服务端口:8000")                 //控制台输出信息
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}