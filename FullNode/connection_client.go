package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func send_tx(ip string, tx *Transaction) {

	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("Tcp dial fail @send_tx:", err)
		return
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	err = encoder.Encode(tx)
	if err != nil {
		fmt.Println("Encode error@send_tx:", err)
		return
	}

	var response string
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Decode error @send_tx:", err)
		return
	}

	fmt.Println("Response from peer:", response)
}
