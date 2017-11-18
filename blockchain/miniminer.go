package blockchain

import (
	"github.com/mcyprian/pivcoin/uuid"
)

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
