package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction //存储交易数据，不再是字符串数据了
	PrevBlockHash []byte
	Nonce         int
	Hash          []byte
	Ip            []string
}

// NewBlock
// Multiple transactions in one block
func NewBlock(transactions []*Transaction, prevBlockHash []byte, electors []string) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, 0, []byte{}, electors}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

// HashTransactions caculate transaction hash

//Get every tx hash, connect every tx, generate a final hash
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// NewGenesisBlock
// create genesis block, need mining when create genesis
func NewGenesisBlock(coninbase *Transaction) *Block {
	return NewBlock([]*Transaction{coninbase}, []byte{}, []string{fullNode1})
}

// Serialize

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
