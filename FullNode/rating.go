package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type rating struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	ID       string `json:"ID"`
	Polarity string `json:"polarity"`
}

func recvRating() {
	listener, err := net.Listen("tcp", ":9884")
	if err != nil {
		fmt.Println("Can't listen to 9884:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server up listen to:", 9884)

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		// 处理客户端连接
		go handleRating(conn)
	}
}

func handleRating(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Reading data error:", err)

	}

	data := buffer[:n]

	fmt.Printf("Received JSON data： %+v\n", data)

	var jsonData rating
	err = json.Unmarshal(data, &jsonData)
	if err != nil {

		fmt.Println("json decoding error")

	}

	sender := jsonData.Sender
	receiver := jsonData.Receiver
	id := jsonData.ID
	polarity, _ := strconv.Atoi(jsonData.Polarity)

	submitRating, ratingReward := ratings(sender, receiver, polarity)

	targetUrl := "http://repustation.000webhostapp.com/index.php"

	postData := map[string]interface{}{
		"id":     id,
		"rating": submitRating,
	}

	jsonPostData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("JSON decode error:", err)

	}
	resp, err := http.Post(targetUrl, "application/json", bytes.NewBuffer(jsonPostData))

	if err != nil {
		fmt.Println("HTTP POST error:", err)
		//return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful!")
	} else {
		fmt.Println("Request failed. Status code:", resp.StatusCode)
	}

	senderReward := coinbaseReward(sender, ratingReward, "")

	invest := getInvest(id)
	reviewerReward := reviewReward(submitRating, invest)

	bc := NewBlockchain()
	currentIP := string(getIPV4())
	punish := punishment(receiver)

	if polarity > 0 {

		receiverReward := coinbaseReward(receiver, reviewerReward+invest, "")
		bc.MineBlock([]*Transaction{senderReward, receiverReward}, []string{currentIP})
	} else if polarity < 0 {

		reviewerGetReward := float64(invest) - (float64(reviewerReward) * (1 + punish))
		if reviewerGetReward > 0 {
			receiverReward := coinbaseReward(receiver, int(reviewerGetReward), "")
			bc.MineBlock([]*Transaction{senderReward, receiverReward}, []string{currentIP})
		} else {
			bc.MineBlock([]*Transaction{senderReward}, []string{currentIP})
		}

	} else {
		fmt.Println("Polarity error")
	}

	bc.Db.Close()
}

func getInvest(id string) int {
	resp, err := http.Get("https://repustation.000webhostapp.com/index.php?ID=" + id)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	if response.Error != "" {
		fmt.Println("Server response error:", response.Error)

	}

	investValue := response.Investment
	return investValue
}

type Response struct {
	Investment int    `json:"investment"`
	Error      string `json:"error"`
}
