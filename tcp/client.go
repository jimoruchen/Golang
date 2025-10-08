package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var text string
		fmt.Printf("请输入内容：")
		fmt.Scanln(&text)
		if text == "exit" {
			conn.Close()
			break
		}
		conn.Write([]byte(text))
	}
}
