package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

type dataJSON struct {
	Sender     string `json:"sender"`
	Seller     string `json:"seller"`
	Comment    string `json:"comment"`
	Investment int    `json:"investment"`
	Polarity   int    `json:"polarity"`
	Ratings    int    `json:"ratings"`
}

func recvReview() {

	listener, err := net.Listen("tcp", ":9887")
	if err != nil {
		fmt.Println("Can't creat connection:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening to :", 9887)

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Reading data error:", err)
		//return
	}

	data := buffer[:n]

	fmt.Printf("Received JSON data： %+v\n", data)

	var jsonData dataJSON
	err = json.Unmarshal(data, &jsonData)
	if err != nil {

		fmt.Println("json decode error")
		return
	}
	seller := jsonData.Seller
	comment := jsonData.Comment
	//sender :=jsonData.Sender
	//investment :=jsonData.Investment
	hashReview(seller, comment)
	//Send http request
	url := "http://repustation.000webhostapp.com/index.php"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("HTTP POST error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful!")
	} else {
		fmt.Println("Request failed. Status code:", resp.StatusCode)
	}
	//return sender, investment
}

func hashReview(seller string, comment string) {
	dbPath := "./reviewHash/" + seller + ".txt"
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {

		existingContent, err := ioutil.ReadFile(dbPath)
		if err != nil {
			fmt.Println("Can't read file:", err)
			return
		}

		updatedContent := string(existingContent) + comment

		hash := sha256.Sum256([]byte(updatedContent))

		err = ioutil.WriteFile(dbPath, []byte(hex.EncodeToString(hash[:])), os.ModePerm)
		if err != nil {
			fmt.Println("Can't write file:", err)
			return
		}

		fmt.Println("File created and insert new hash value ")

	} else {

		hash := sha256.Sum256([]byte(comment))

		err := ioutil.WriteFile(dbPath, []byte(hex.EncodeToString(hash[:])), os.ModePerm)
		if err != nil {
			fmt.Println("Unable to create and insert hash value:", err)
			return
		}
		fmt.Println("File created and insert new hash value")
	}

}
