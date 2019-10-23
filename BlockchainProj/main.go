package main

import (
	"fmt"

	blockchainCore "./blockchaincore"
)

func main() {

	//https://jeiwan.cc/posts/building-blockchain-in-go-part-1/
	bc := blockchainCore.NewBlockchain()

	bc.AddBlock("Send 1 BTC to pino paperino")
	bc.AddBlock("Send 2 more BTC to pino paperino")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
