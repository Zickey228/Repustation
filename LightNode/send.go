package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func send(from, to string, amount int) {
	//发送代币给指定地址
	if !ValidateAddress(from) {
		log.Panic("ERROR: 发送地址非法")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: 接收地址非法")
	}
	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库

	tx := NewUTXOTransaction(from, to, amount, bc) //创建交易
	send_tx(fullnode1, tx)
}

func send_tx(ip string, tx *Transaction) {
	// 建立TCP连接
	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("无法建立TCP连接:", err)
		return
	}
	defer conn.Close()

	// 创建一个编码器和解码器
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	// 创建一个结构体实例

	// 发送结构体到服务器
	err = encoder.Encode(tx)
	if err != nil {
		fmt.Println("编码错误:", err)
		return
	}

	// 接收服务器的响应
	var response string
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	fmt.Println("服务器响应:", response)
}
