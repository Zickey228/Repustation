package main

import (
	"log"
)

func listWallets() []string {
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()
	return addresses
	//for _, address := range addresses {
	//	fmt.Println(address)
	//}
}
