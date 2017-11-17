package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/mcyprian/pivcoin/uuid"
)

func intToBytes(num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 31415926)
	return bs
}

func NewGenesisBlock() *Block {
	block := &Block{
		Index:        0,
		PreviousHash: []byte{0},
		Timestamp:    uint32(time.Now().Unix()),
		Data: Data{
			uuid.GenerateID(),
			uuid.GenerateID(),
			10000},
	}
	block.CountHashSum()
	return block
}

type Data struct {
	FromAddress []byte
	ToAddress   []byte
	Amount      uint32
}

type Block struct {
	Index        uint32
	PreviousHash []byte
	Timestamp    uint32
	// TODO add nonce to identify proof-of-work step
	Data Data
	Hash []byte
}

func (block *Block) CountHashSum() {
	h := sha256.New()
	bs := []byte{}
	bs = append(bs, intToBytes(block.Index)...)
	bs = append(bs, block.PreviousHash...)
	bs = append(bs, intToBytes(block.Timestamp)...)
	bs = append(bs, block.Data.FromAddress...)
	bs = append(bs, block.Data.ToAddress...)
	bs = append(bs, intToBytes(block.Data.Amount)...)
	block.Hash = h.Sum(bs)
}

func (block *Block) GetHashSum() []byte {
	h := sha256.New()
	bs := []byte{}
	bs = append(bs, intToBytes(block.Index)...)
	bs = append(bs, block.PreviousHash...)
	bs = append(bs, intToBytes(block.Timestamp)...)
	bs = append(bs, block.Data.FromAddress...)
	bs = append(bs, block.Data.ToAddress...)
	bs = append(bs, intToBytes(block.Data.Amount)...)
	return h.Sum(bs)
}

func (block Block) ToJSON() string {
	repr, err := json.MarshalIndent(block, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(repr)
}
