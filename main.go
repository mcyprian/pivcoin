package main

import (
	"fmt"

	"github.com/mcyprian/pivcoin/blockchain"
)

const (
	// Harcoded JSON of genesis block
	blockData = `
{
    "Index": 0,
    "PreviousHash": "AA==",
    "Timestamp": 1511008886,
    "Nonce": 0,
    "Data": [
        {
            "FromAddress": "Ja7R69CIXq1AtGUcXPGmHg==",
            "ToAddress": "BOwrfUTBBtBPnkA1qXDtfg==",
            "Amount": 10000
        }
    ],
    "Hash": "dl7fAQB2Xt8Bdl7fASWu0evQiF6tQLRlHFzxph4E7Ct9RMEG0E+eQDWpcO1+dl7fAeOwxEKY/BwUmvv0yJlvuSQnrkHkZJuTTKSVmRt4UrhV"
}`
)

func main() {
	genesis := blockchain.NewFromJSON([]byte(blockData))
	fmt.Println("Genesis block:")
	fmt.Println(genesis.ToJSON())

	fmt.Println("Following valid block found!:")
	fmt.Println(blockchain.FindNextBlock(genesis).ToJSON())
}
