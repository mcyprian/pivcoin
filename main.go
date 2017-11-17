package main

import (
	"fmt"
	"github.com/mcyprian/pivcoin/blockchain"
)

func main() {
	block := blockchain.NewGenesisBlock()
	fmt.Println(block)
	fmt.Println("\n\n")
	fmt.Println(block.ToJSON())
}
