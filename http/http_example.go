package main

import (
	"fmt"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.parseFrom()       //解析参数，默认是不解析的
	fmt.Println(r.Form) //这些是输出到服务器打印的log日志
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println()

}