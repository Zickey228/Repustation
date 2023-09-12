package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkReview() {

	listener, err := net.Listen("tcp", ":9886")
	if err != nil {
		fmt.Println("Can't listening 9886:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server up:", 9886)

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Peer connection error:", err)
			continue
		}

		go handleCheck(conn)
	}
}

func handleCheck(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Reading data error:", err)
		return
	}
	target := string(buffer[:n])

	fmt.Printf("Server is checking the hash of seller : %s", target)
	dbPath := "./reviewHash/" + target + ".txt"
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {

		existingContent, err := ioutil.ReadFile(dbPath)
		_, err = conn.Write(existingContent)
		if err != nil {
			fmt.Println("Error sending response:", err.Error())
			return
		}
	} else {
		fmt.Println("File not exist")

	}

}
