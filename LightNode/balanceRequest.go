package main

import (
	"fmt"
	"net"
)

// dns 服务器ip

func balanceReq(address string) string {
	conn, err := net.Dial("tcp", dns+":9888")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())

	}
	defer conn.Close()

	// 发送数据到服务器
	message := "bal"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())

	}

	_, err = conn.Write([]byte(address))
	if err != nil {
		fmt.Println("Error sending string data: %s", err)
	}

	// 读取服务器的响应
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err.Error())

	}

	response := string(buffer)
	fmt.Println("Response from server:", response)
	return response
}
