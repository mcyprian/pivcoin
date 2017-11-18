package blockchain

import (
	"encoding/json"
	"fmt"
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

func BuildChain(length int) []*Block {
	var previous *Block
	genesis := NewFromJSON([]byte(blockData))
	fmt.Println("Genesis block:")
	fmt.Println(genesis.ToJSON())
	chain := []*Block{genesis}
	previous = genesis

	fmt.Println("Building blockchain:")
	for i := 0; i < length; i++ {
		newblock := FindNextBlock(previous)
		chain = append(chain, newblock)
		fmt.Println(newblock.ToJSON())
		fmt.Println(newblock.Hash)
		previous = newblock
	}
	return chain
}

func SerializeChain(chain []*Block) string {
	repr, err := json.MarshalIndent(chain, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(repr)
}
