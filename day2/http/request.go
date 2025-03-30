package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 写一个简单的http
func main() {
	// http服务
	// 监听端口
	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
	// 绑定一个get请求,路径为home
	http.HandleFunc("/home", handler)
	if err != nil {
		fmt.Println("http server start error")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	req, _ := io.ReadAll(body)
	// 确保请求体在函数结束时关闭
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			fmt.Println("body close error")
		}
	}(body)
	fmt.Printf("request body: %s\n", string(req))
	query := r.URL.Query()
	q := query["q"][0]
	fmt.Printf("q: %s\n", q)
	marshal, _ := json.Marshal(r.URL)
	// 把marshal输出为json格式
	println(string(marshal))
	r.ParseForm()
	// 响应
	w.Write([]byte("hello world"))
}
