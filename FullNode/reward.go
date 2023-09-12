package main

import "fmt"

func coinbaseReward(to string, amount int, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to :%s", to)
	}

	txin := TxInput{[]byte{}, -1, nil, []byte(data)}
	txout := NewTxOutput(amount, to)
	tx := Transaction{nil, []TxInput{txin}, []TxOutput{*txout}}
	tx.ID = tx.Hash()

	return &tx
}
