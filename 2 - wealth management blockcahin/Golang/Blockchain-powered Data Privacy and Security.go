package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Data     string
	Previous *Block
}

type Blockchain struct {
	Head *Block
}

func NewBlock(data string, previous *Block) *Block {
	block := &Block{
		Data:     data,
		Previous: previous,
	}
	return block
}

func (b *Block) CalculateHash() string {
	hash := sha256.Sum256([]byte(b.Data))
	return hex.EncodeToString(hash[:])
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	genesisBlock := NewBlock("Genesis Block", nil)
	blockchain.Head = genesisBlock
	return blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	previousHash := bc.Head.CalculateHash()
	newBlock := NewBlock(data, bc.Head)
	bc.Head = newBlock
	fmt.Printf("Block added: %s\n", newBlock.Data)
	fmt.Printf("Previous hash: %s\n", previousHash)
	fmt.Printf("New block hash: %s\n", newBlock.CalculateHash())
}

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("Client data 1")
	blockchain.AddBlock("Client data 2")
	blockchain.AddBlock("Client data 3")
}
