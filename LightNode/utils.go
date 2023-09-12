package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
)

// IntToHex
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func SaveToTxt(content string, filename string) {
	file := "./" + filename + ".txt"
	newFile, err := os.Create(file)
	if err != nil {
		log.Panic(err)
		return
	}
	defer newFile.Close()
	newFile.WriteString(content)
}

func LoadFileToHex(fileName string) ([]byte, error) {
	newFile := "./" + fileName + ".txt"
	data, err := ioutil.ReadFile(newFile)
	//str := string(data)
	//str, _ := new(big.Int).SetString(string(data), 10)
	return data, err
}
