package main

import (
	"fmt"

	"github.com/yourusername/blockchain/easyblockchain"
)

func main() {
	bc := easyblockchain.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to derick")
	bc.AddBlock("Send 1 more BTC to cosmos")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		pow := easyblockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %t\n", pow.Validate())
		fmt.Println()
	}
}
