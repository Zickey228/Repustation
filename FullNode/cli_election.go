package main

import "fmt"

func (cli *CLI) election(myAddress string) {

	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()
	lastBlock := bci.Next()
	target := lastBlock.Ip

	myIp := string(getIPV4())

	sendElection(bc, myIp, myAddress, target[0])
	fmt.Printf("Sending election request to %s,My ip：%s,My address：%s", target[0], myIp, myAddress)
}
