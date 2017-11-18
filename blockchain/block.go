package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"reflect"
	"time"

	"github.com/mcyprian/pivcoin/uuid"
)

func intToBytes(num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 31415926)
	return bs
}

// Constructor of genesis block
func NewGenesisBlock() *Block {
	block := &Block{
		Index:        0,
		PreviousHash: []byte{0},
		Timestamp:    uint32(time.Now().Unix()),
		Nonce:        0,
		Data: []Transaction{
			Transaction{
				uuid.GenerateID(),
				uuid.GenerateID(),
				10000},
		},
	}
	block.CountHashSum()
	return block
}

// Create new Block object from serialized JSON data
func NewFromJSON(data []byte) *Block {
	var block Block
	if err := json.Unmarshal(data, &block); err != nil {
		panic(err)
	}
	return &block
}

// Create new Block object based on predecessor in the chain
func NewFromPrevious(previous *Block, nonce uint32, data []Transaction) *Block {
	block := &Block{
		Index:        previous.Index + 1,
		PreviousHash: previous.Hash,
		Timestamp:    uint32(time.Now().Unix()),
		Data:         data,
		Nonce:        nonce,
	}
	block.CountHashSum()
	return block
}

type Transaction struct {
	FromAddress []byte
	ToAddress   []byte
	Amount      uint32
}

func (data *Transaction) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, data.FromAddress...)
	bs = append(bs, data.ToAddress...)
	bs = append(bs, intToBytes(data.Amount)...)
	return bs
}

// The main Block structure
type Block struct {
	Index        uint32
	PreviousHash []byte
	Timestamp    uint32
	Nonce        uint32
	Data         []Transaction
	Hash         []byte
}

func (block *Block) CountHashSum() {
	block.Hash = block.GetHashSum()
}

// Calculate hash sum of block attributes
func (block *Block) GetHashSum() []byte {
	h := sha256.New()
	bs := []byte{}
	bs = append(bs, intToBytes(block.Index)...)
	bs = append(bs, block.PreviousHash...)
	bs = append(bs, intToBytes(block.Timestamp)...)
	bs = append(bs, intToBytes(block.Nonce)...)
	for _, item := range block.Data {
		bs = append(bs, item.Serialize()...)
	}
	h.Write(bs)
	return h.Sum(nil)
}

// Serialize block to JSON
func (block Block) ToJSON() string {
	repr, err := json.MarshalIndent(block, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(repr)
}

// Determine whether block is a valid successor of previous block of not
// TODO it is your job to finish this Risko ;)
func (block *Block) IsNext(previous *Block) bool {
	if block.Index != previous.Index+1 {
		return false
	} else if !reflect.DeepEqual(block.PreviousHash, previous.Hash) {
		return false
	} else if block.Timestamp <= previous.Timestamp {
		return false
	} else if block.Hash[0] != 0 {
		return false
	}
	return true
}
