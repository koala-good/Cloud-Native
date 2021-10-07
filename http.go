package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/", myResponse)
	http.ListenAndServe("127.0.0.1:8080", nil)
	fmt.Println(runtime.Version())

}

func myResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您输入的地址:", r.URL.Path)
	fmt.Fprintln(w, "字符串:", r.URL.RawQuery)
	fmt.Fprintln(w, "头部信息全部为", r.Header)
	fmt.Fprintln(w, "头部信息中[Accept]的信息为:", r.Header["Accept"])

}
