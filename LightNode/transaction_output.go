package main

import "bytes"

type TxOutput struct {
	Value int

	PubKeyHash []byte
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func NewTxOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (out *TxOutput) Lock(address []byte) {
	expubKeyHash := Base58Decode(address)
	pubKeyHash := expubKeyHash[1 : len(expubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}
