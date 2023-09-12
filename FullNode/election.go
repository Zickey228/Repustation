package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
)

type WeightedData struct {
	Data   string
	Weight int
}

type ElectionRequest struct {
	Ip      string
	Address string
}

func sendElection(bc *Blockchain, myIp string, myAddress string, targetIp string) {

	conn, err := net.Dial("tcp", targetIp+":9885")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	request := ElectionRequest{
		Ip:      myIp,
		Address: myAddress,
	}
	bal := getElectorBalance(bc, "1AkFuweFVhr4pVkGnoua16qWtJpUb8NhZh")
	fmt.Println("The weight of this address is ：", bal)

	requestJSON, err := json.Marshal(request)
	fmt.Println(requestJSON)
	if err != nil {
		fmt.Println("Error encoding request:", err.Error())
		return
	}

	_, err = conn.Write(requestJSON)
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}
	fmt.Println("Election request sent")

}

var electionList = []WeightedData{}

func recvElection(bc *Blockchain, nonce int) []string {

	electionList = []WeightedData{}

	listener, err := net.Listen("tcp", ":9885")
	if err != nil {
		fmt.Println("Error listening:", err.Error())

	}
	defer listener.Close()
	fmt.Println("Server is listening on 9885 to receive election request")

	for i := 0; i < electSize; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleElection(bc, conn)
	}

	fmt.Println("Finish collect election list", electionList)
	electors := random(int64(nonce), electionList, selectSize)
	fmt.Println(electors)
	return electors
}

func handleElection(bc *Blockchain, conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	//fmt.Println("Unmarshalled data is :", buffer)

	var request ElectionRequest
	err = json.Unmarshal(buffer[:n], &request)
	if err != nil {
		fmt.Println("Error decoding JSON:", err.Error())
		return
	}
	//fmt.Println("Marshalled data is：", request)

	weight := getElectorBalance(bc, request.Address) + 1

	fmt.Println("Request address is：", request.Address)
	fmt.Println("The weight of this address is：", weight)

	data := WeightedData{
		Data:   request.Ip,
		Weight: weight,
	}

	electionList = append(electionList, data)

}

func random(randomSeed int64, data []WeightedData, n int) []string {

	rand.Seed(randomSeed)
	totalWeight := 0
	for _, data := range data {
		totalWeight += data.Weight
	}

	if totalWeight == 0 {
		fmt.Println("No valid data with positive weight.")
		return []string{}
	}

	selectedData := map[string]bool{}

	for len(selectedData) < n {

		randomNumber := rand.Intn(totalWeight)

		for _, data := range data {
			if randomNumber < data.Weight && !selectedData[data.Data] {
				selectedData[data.Data] = true
				fmt.Printf("Selected data: %s\n", data.Data)
				break
			}
			randomNumber -= data.Weight
		}
	}

	electors := []string{}
	for elector := range selectedData {
		electors = append(electors, elector)
	}
	return electors
}

func getElectorBalance(bc *Blockchain, address string) int {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Illegal address")
	}

	balance := 0
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("'%s' balance is %d\n", address, balance)
	return balance
}
