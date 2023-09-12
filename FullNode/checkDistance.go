package main

import (
	"errors"
)

func checkDistance(senderAddress string, receiverAddress string) (int, error) {
	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()
	var lastBlockHeight int
	foundTransaction := false

	for {
		block := bci.Next()
		lastBlockHeight++

		for _, tx := range block.Transactions {

			if ContainsAddress(tx, senderAddress) && ContainsAddress(tx, receiverAddress) {
				foundTransaction = true
				break
			}
		}

		if foundTransaction {
			return lastBlockHeight, nil
		}

		if len(block.PrevBlockHash) == 0 {
			return 0, errors.New("未找到包含指定交易的区块")
		}
	}
}
func checkBigDistance(receiverAddress string, threshhold int) (int, error) {
	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()
	var lastBlockHeight int
	foundTransaction := false

	for {
		block := bci.Next()
		lastBlockHeight++

		for _, tx := range block.Transactions {

			for i := 0; i < len(tx.Vout); i++ {
				if ContainsBigAdd(tx, receiverAddress) && tx.Vout[0].Value > threshhold {
					foundTransaction = true
					break
				}
			}

		}

		if foundTransaction {
			return lastBlockHeight, nil
		}

		if len(block.PrevBlockHash) == 0 {
			return 0, errors.New("未找到包含指定交易的区块")
		}
	}
}

func ContainsAddress(tx *Transaction, address string) bool {

	for _, vin := range tx.Vin {
		if vin.UsesKey([]byte(address)) {
			return true
		}
	}

	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	for _, vout := range tx.Vout {
		if vout.IsLockedWithKey(pubKeyHash) {
			return true
		}
	}
	return false
}

func ContainsBigAdd(tx *Transaction, address string) bool {

	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	for _, vout := range tx.Vout {
		if vout.IsLockedWithKey(pubKeyHash) {
			return true
		}
	}
	return false
}
