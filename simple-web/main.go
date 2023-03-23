package main

import "net/http"

func main() {
	//监听的地址
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// 监听本地8080端口
	http.ListenAndServe("localhost:8080", nil) // default serverMux
}
