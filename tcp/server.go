package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("listening on", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端：" + conn.RemoteAddr().String() + "连接...")
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Println("客户端" + conn.RemoteAddr().String() + "断开连接...")
				break
			}
			fmt.Println(string(buf[:n]))
		}
	}
}
