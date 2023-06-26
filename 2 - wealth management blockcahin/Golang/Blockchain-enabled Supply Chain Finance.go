package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Index     int
	Timestamp string
	Data      interface{}
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain.
type Blockchain struct {
	Chain []Block
}

// CalculateHash calculates the hash for a block.
func CalculateHash(index int, timestamp string, data interface{}, prevHash string) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%v%s", index, timestamp, data, prevHash)))
	return hex.EncodeToString(hash.Sum(nil))
}

// CreateGenesisBlock creates the genesis block of the blockchain.
func CreateGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}
}

// CreateNewBlock creates a new block in the blockchain.
func CreateNewBlock(index int, data interface{}, prevHash string) Block {
	timestamp := time.Now().String()
	hash := CalculateHash(index, timestamp, data, prevHash)

	return Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      hash,
	}
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(data interface{}) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := CreateNewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

// PrintBlockchain prints the entire blockchain.
func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Chain {
		blockJSON, _ := json.MarshalIndent(block, "", "  ")
		fmt.Println(string(blockJSON))
		fmt.Println("------------------------")
	}
}

func main() {
	// Create the genesis block
	genesisBlock := CreateGenesisBlock()

	// Create the blockchain
	blockchain := Blockchain{
		Chain: []Block{genesisBlock},
	}

	// Add blocks to the blockchain
	blockchain.AddBlock("Transaction 1")
	blockchain.AddBlock("Transaction 2")
	blockchain.AddBlock("Transaction 3")

	// Print the blockchain
	blockchain.PrintBlockchain()
}
