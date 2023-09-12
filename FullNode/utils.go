package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func getIPV4() []byte {
	resp, err := http.Get("https://ipv4.netarm.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, _ := ioutil.ReadAll(resp.Body)
	return ip
}
