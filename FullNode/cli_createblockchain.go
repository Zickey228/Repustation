package main

import (
	"fmt"
	"log"
)

// createBlockchain
func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Illegal address")
	}
	bc := CreatBlockchain(address)
	bc.Db.Close()
	fmt.Println("New blockchain created！")
}
