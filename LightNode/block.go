package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction //存储交易数据，不再是字符串数据了
	PrevBlockHash []byte
	Nonce         int
	Hash          []byte
	Ip            []string
}

func DeserializeBlock(d []byte) *Block {
	var block Block //一般都不会通过指针来创建一个struct。记住struct是一个值类型

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block //返回block的引用
}
