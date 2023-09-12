package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handler_sendFile(conn net.Conn) error {
	file, err := os.Open(dbFile)
	if err != nil {
		return fmt.Errorf("Error opening Bolt file: %s", err)
	}
	defer file.Close()

	_, err = conn.Write([]byte("upd"))
	if err != nil {
		return fmt.Errorf("Error sending status code: %s", err)
	}

	_, err = io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("Error sending Bolt file: %s", err)
	}

	return nil
}

func send_file(ip string) {

	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("Can't create connection:", err)
		return
	}
	defer conn.Close()

	err = handler_sendFile(conn)
	if err != nil {
		fmt.Println("Error sending Bolt file:", err)
		return
	}

	fmt.Println("Send file successfully")
}
