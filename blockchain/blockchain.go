package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	STORAGEPATH = "./blockchain_replica"
)

// Temprary function to simulate blockchain creation
func BuildChain(length int) []*Block {
	handler, err := NewHandlerJSON()
	if err != nil {
		var previous *Block
		genesis := CountGenesisBlock()
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
		handler := HandlerJSON{chain}
		handler.Commit()
	}
	data, err := handler.Serialize()
	if err != nil {
		fmt.Println("ERROR: Failed to serialize the chain")
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	return handler.chain
}

// API to work with blockchain
type BlockchainHandler interface {
	Init() error
	InsertNew(block *Block) error
	GetLast() *Block
	GetByHash(hash []byte) *Block
	GetByIndex(hash []byte) *Block
	GetAll() []Block
	Commit() error
}

type HandlerJSON struct {
	chain []*Block
}

func NewHandlerJSON() (*HandlerJSON, error) {
	handler := HandlerJSON{}
	if _, err := os.Stat(STORAGEPATH); !os.IsNotExist(err) {
		ierr := handler.Init()
		if ierr != nil {
			return nil, ierr
		}
		return &handler, nil
	}
	fmt.Println("WARNINIG: Handler contains empty blockchain")
	return &handler, nil
}

func (handler *HandlerJSON) Serialize() ([]byte, error) {
	data, err := json.MarshalIndent(handler.chain, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (handler *HandlerJSON) Commit() error {
	data, err := handler.Serialize()
	if err != nil {
		return err
	}
	werr := ioutil.WriteFile(STORAGEPATH, data, 0644)
	if werr != nil {
		return werr
	}
	return nil
}

func (handler *HandlerJSON) Init() error {
	data, err := ioutil.ReadFile(STORAGEPATH)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &handler.chain); err != nil {
		return err
	}
	return nil
}
