package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("   createblockchain -address ADDRESS ~ Create new blockchain and get reward to this address")
	fmt.Println("   createwallet ~ create new key pair and store to wallet file")
	fmt.Println("   getbalance -address ADDRESS  ~ get address balance")
	fmt.Println("   listaddresses ~ list all the addresses in wallet file")
	fmt.Println("   printchain ~ print all the blocks in blockchain file")
	fmt.Println("   send -from FROM -to To -amount ~ send <amount> tokens from <from> to <to>")
	fmt.Println("   pack -address ADDR ~ Use this address to pack transaction ")
	fmt.Println("   sync ~ synchronise blockchain from peer")
	fmt.Println("   elect -address ADDR ~ use this address to elect ")
	fmt.Println("   recvRating ~ use address to receive ratings from other peers")
	fmt.Println("   recvReview ~ receive reviews from other peers")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	packCmd := flag.NewFlagSet("packTxs", flag.ExitOnError)
	syncCmd := flag.NewFlagSet("sync", flag.ExitOnError)
	electCmd := flag.NewFlagSet("elect", flag.ExitOnError)
	recvRatingCmd := flag.NewFlagSet("recvRating", flag.ExitOnError)
	recvReviewCmd := flag.NewFlagSet("recvReview", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "Address receiving tokens")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "Receiving genesis rewards")
	sendFrom := sendCmd.String("from", "", "wallet address")
	sendTo := sendCmd.String("to", "", "target wallet address")
	sendAmount := sendCmd.Int("amount", 0, "transfer amount")
	packAddress := packCmd.String("address", "", "pack address")
	electAddress := electCmd.String("address", "", "elect address")

	switch os.Args[1] {
	case "getbalance":

		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":

		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "pack":

		err := packCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "sync":
		err := syncCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "elect":
		err := electCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "recvRating":
		err := recvRatingCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "recvReview":
		err := recvReviewCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createWalletCmd.Parsed() {
		cli.createWallet()
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if packCmd.Parsed() {
		if *packAddress == "" {
			packCmd.Usage()
			os.Exit(1)
		}
		cli.pack(*packAddress)
	}

	if syncCmd.Parsed() {
		cli.sync()
	}

	if packCmd.Parsed() {
		if *packAddress == "" {
			packCmd.Usage()
			os.Exit(1)
		}
		cli.pack(*packAddress)
	}

	if electCmd.Parsed() {
		if *electAddress == "" {
			electCmd.Usage()
			os.Exit(1)
		}
		cli.election(*electAddress)
	}

	if recvRatingCmd.Parsed() {
		cli.recvRating()
	}

	if recvReviewCmd.Parsed() {
		cli.recvReview()
	}
}
