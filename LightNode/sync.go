package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const dbFile = "blockchain.db"

// 从全节点请求数据
func recv_file(ip string) {
	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	err = handler_recv(conn)
	if err != nil {
		fmt.Println("Error receiving Bolt file:", err)
		return
	}

	fmt.Println("Client finished execution")
}

func handler_recv(conn net.Conn) error {
	file, err := os.Create(dbFile)
	if err != nil {
		return fmt.Errorf("Error creating file: %s", err)
	}
	defer file.Close()

	_, err = conn.Write([]byte("syn"))
	if err != nil {
		return fmt.Errorf("Error sending status code: %s", err)
	}

	_, err = io.Copy(file, conn)
	if err != nil {
		return fmt.Errorf("Error receiving Bolt file: %s", err)
	}

	return nil
}
