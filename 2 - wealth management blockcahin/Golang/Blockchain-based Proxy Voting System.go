package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Vote      string
	Hash      string
	PrevHash  string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Chain []Block
}

// CreateBlock creates a new block in the blockchain
func (bc *Blockchain) CreateBlock(index int, timestamp, vote, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: timestamp,
		Vote:      vote,
		PrevHash:  prevHash,
	}

	block.Hash = block.CalculateHash()
	return block
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	data := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Vote, b.PrevHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(block Block) {
	bc.Chain = append(bc.Chain, block)
}

// VerifyBlock verifies the integrity of a block
func (b *Block) VerifyBlock() bool {
	return b.Hash == b.CalculateHash()
}

// VerifyChain verifies the integrity of the entire blockchain
func (bc *Blockchain) VerifyChain() bool {
	chain := bc.Chain

	for i := 1; i < len(chain); i++ {
		currentBlock := chain[i]
		previousBlock := chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

func main() {
	// Create a new blockchain
	blockchain := Blockchain{}

	// Create and add blocks to the blockchain
	block1 := blockchain.CreateBlock(1, "2023-06-26 10:00:00", "Vote A", "")
	block2 := blockchain.CreateBlock(2, "2023-06-26 11:00:00", "Vote B", block1.Hash)
	block3 := blockchain.CreateBlock(3, "2023-06-26 12:00:00", "Vote C", block2.Hash)

	blockchain.AddBlock(block1)
	blockchain.AddBlock(block2)
	blockchain.AddBlock(block3)

	// Verify the integrity of the blockchain
	fmt.Println("Blockchain is valid:", blockchain.VerifyChain())
}
