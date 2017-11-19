package blockchain

import (
	"github.com/mcyprian/pivcoin/uuid"
)

const (
	// Harcoded JSON of genesis block
	blockData = `
{
    "Index": 0,
    "PreviousHash": "AA==",
    "Timestamp": 1511008886,
    "Nonce": 0,
    "Data": [ { } ],
    "Hash": "AA=="
}`
)

func CountGenesisBlock() *Block {
	genesis := NewFromJSON([]byte(blockData))
	var nonce uint32 = 0
	for {
		genesis.Nonce = nonce
		genesis.CountHashSum()
		if genesis.Hash[0] == 0 && genesis.Hash[1] == 0 {
			return genesis
		}
		nonce++
	}
}

func FindNextBlock(previous *Block) *Block {
	var nonce uint32 = 0
	for {
		block := NewFromPrevious(previous, nonce, []Transaction{
			Transaction{
				uuid.GenerateID(),
				uuid.GenerateID(),
				10000},
		})
		if block.IsNext(previous) {
			return block
		}
		nonce++
	}
}
