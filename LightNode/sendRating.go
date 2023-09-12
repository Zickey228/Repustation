package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type rating struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	ID       string `json:"ID"`
	Polarity string `json:"polarity"`
}

func sendRating(sender string, receiver string, id string, polarity string) {
	conn, err := net.Dial("tcp", ip2+":9884")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
	}
	defer conn.Close()

	// JSON 格式的数据
	data := rating{
		Sender:   sender,
		Receiver: receiver,
		ID:       id,
		Polarity: polarity,
	}

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encode error:", err)
		return
	}

	_, err = conn.Write(json)
	if err != nil {
		fmt.Println("Send json error:", err)
		return
	}

	fmt.Println("json sent to server")

}
