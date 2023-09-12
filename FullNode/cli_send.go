package main

import (
	"fmt"
	"log"
)

// send 转账
func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sending address illegal")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Receiving address illegal")
	}

	bc := NewBlockchain()
	defer bc.Db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)

	//最后一个区块
	bci := bc.Iterator()
	lastBlock := bci.Next()
	for _, ip := range lastBlock.Ip {
		currentIP := string(getIPV4())
		fmt.Printf("current ip : %s, target ip :%s", currentIP, ip)
		if ip == currentIP {
			fmt.Println("Current address in use is miner address")
			reward := NewCoinbaseTX(from, "")
			bc.MineBlock([]*Transaction{reward, tx}, []string{currentIP})

			send_file(dns)

		} else { //发送给所有挖矿节点

			send_tx(ip, tx)
			fmt.Printf("Sending data to ----->" + ip)
		}

	}

	fmt.Println("Sending tokens successfully")
}
