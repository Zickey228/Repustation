package main

const dns = "3.10.214.121"
const fullNode1 = "18.134.137.142"
const fullNode2 = "13.40.136.108"
const fullNode3 = "13.40.136.108"

const blockSize = 1
const selectSize = 1
const electSize = 3

const disThreshold = 2
const roi = 1 / 10
const punishThreshold = 20

func main() {
	cli := CLI{}
	cli.Run()

}
