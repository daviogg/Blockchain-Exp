package main

import (
	"fmt"

	functions "BlockchainProj/functions"
)

//#region - types section

func main() {
	bc := functions.NewBlockchain()

	bc.AddBlock("Send 1 BTC to pino paperino")
	bc.AddBlock("Send 2 more BTC to pino paperino")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
