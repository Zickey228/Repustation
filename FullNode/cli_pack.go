package main

import (
	"fmt"
	"log"
)

func (cli *CLI) pack(addr string) {
	if !ValidateAddress(addr) {
		log.Panic("ERROR: Illegal address when packing")
	}

	bc := NewBlockchain()
	defer bc.Db.Close()

	bci := bc.Iterator()
	lastBlock := bci.Next()
	nonce := lastBlock.Nonce
	fmt.Println("nonce value is:", nonce)

	txs := recv_tx(addr)
	fmt.Println(txs)
	electors := recvElection(bc, nonce)
	fmt.Println(electors)
	bc.MineBlock(txs, electors)

}
