package main

import (
	"fmt"
	"io"
	"net/http"
)

func Index1(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.URL.Path)
	if request.Method != "GET" {
		byteData, _ := io.ReadAll(request.Body)
		fmt.Println(string(byteData))
	}
	fmt.Println(request.Header)
	writer.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/index", Index1)
	fmt.Printf("Server started at port http://127.0.0.1:8080\n")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
