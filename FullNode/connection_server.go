package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

var txs []*Transaction

func recv_tx(address string) []*Transaction {

	txs = []*Transaction{}

	listener, err := net.Listen("tcp", ":9888")
	if err != nil {
		fmt.Println("Can't listen to port @ recv_tx:", err)
	}
	defer listener.Close()

	fmt.Println("Port 9888 up for receiving data")

	for i := 0; i < blockSize; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}

	send_status(dns, "upd")
	send_file(dns)
	fmt.Println("Update to dns succeed")
	return txs
}

func handleConnection(conn net.Conn) {

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	var tx *Transaction
	err := decoder.Decode(&tx)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Println("Received struct body:", tx)

	txs = append(txs, tx)

	response := "Struct body received "
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
}
