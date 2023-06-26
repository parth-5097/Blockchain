package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Asset represents a digital asset
type Asset struct {
	ID          string
	Name        string
	Description string
	Owner       string
	CreatedAt   time.Time
}

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp time.Time
	Asset     Asset
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block in the blockchain
func NewBlock(index int, asset Asset, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now(),
		Asset:     asset,
		PrevHash:  prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

// AddBlock adds a new block to the blockchain
func (chain *Blockchain) AddBlock(asset Asset) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, asset, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// calculateHash calculates the hash of a block
func calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%s%s%s%s", block.Index, block.Timestamp.String(), block.Asset.ID, block.Asset.Owner, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// PrintBlockchain prints the blockchain
func (chain *Blockchain) PrintBlockchain() {
	for _, block := range chain.Blocks {
		fmt.Printf("Block Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp.String())
		fmt.Printf("Asset ID: %s\n", block.Asset.ID)
		fmt.Printf("Asset Name: %s\n", block.Asset.Name)
		fmt.Printf("Asset Description: %s\n", block.Asset.Description)
		fmt.Printf("Asset Owner: %s\n", block.Asset.Owner)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

func main() {
	// Create a new blockchain
	blockchain := Blockchain{
		Blocks: []*Block{},
	}

	// Create a sample asset
	asset := Asset{
		ID:          "123",
		Name:        "Sample Asset",
		Description: "This is a sample asset",
		Owner:       "Alice",
		CreatedAt:   time.Now(),
	}

	// Add the asset to the blockchain
	blockchain.AddBlock(asset)

	// Print the blockchain
	blockchain.PrintBlockchain()

	// Encode the blockchain to JSON
	jsonData, err := json.Marshal(blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON representation of the blockchain
	fmt.Printf("Blockchain JSON:\n%s\n", string(jsonData))
}
