package main

import (
	"fmt"
	"strconv"
)

// printChain
func (cli *CLI) printChain() {
	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. Hash:%x\n", block.PrevBlockHash)

		fmt.Printf("Hash:%x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
