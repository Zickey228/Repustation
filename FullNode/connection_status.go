package main

import (
	"fmt"
	"net"
)

func send_status(ip string, status string) {

	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("Can't connect with peer:", err)
		return
	}
	defer conn.Close()

	message := status

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Sending data error:", err)
		return
	}

	fmt.Println("Status sending successfully")
}
