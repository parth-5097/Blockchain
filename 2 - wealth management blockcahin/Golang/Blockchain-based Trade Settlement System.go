package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Block struct {
	Index        int       `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Data         string    `json:"data"`
	PreviousHash string    `json:"previousHash"`
	Hash         string    `json:"hash"`
	Nonce        int       `json:"nonce"`
}

type Blockchain struct {
	Chain      []Block `json:"chain"`
	Difficulty int     `json:"difficulty"`
}

func calculateHash(index int, timestamp time.Time, data string, previousHash string, nonce int) string {
	hashString := strconv.Itoa(index) + timestamp.String() + data + previousHash + strconv.Itoa(nonce)
	return getSHA256Hash(hashString)
}

func getSHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (b *Block) mineBlock(difficulty int) {
	target := getTargetString(difficulty)
	for {
		b.Nonce++
		b.Hash = calculateHash(b.Index, b.Timestamp, b.Data, b.PreviousHash, b.Nonce)
		if b.Hash[:difficulty] == target {
			break
		}
	}
}

func getTargetString(difficulty int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(difficulty)+"x", 0)
}

func (bc *Blockchain) createGenesisBlock() Block {
	return Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         "Genesis Block",
		PreviousHash: "0",
		Nonce:        0,
	}
}

func (bc *Blockchain) getLatestBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) addBlock(newBlock Block) {
	newBlock.PreviousHash = bc.getLatestBlock().Hash
	newBlock.mineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) isChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != calculateHash(currentBlock.Index, currentBlock.Timestamp, currentBlock.Data, currentBlock.PreviousHash, currentBlock.Nonce) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

var blockchain Blockchain

func handleGetChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain)
}

func handleAddBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type AddBlockRequest struct {
		Data string `json:"data"`
	}

	var request AddBlockRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	newBlock := Block{
		Index:        len(blockchain.Chain),
		Timestamp:    time.Now(),
		Data:         request.Data,
		PreviousHash: "",
		Nonce:        0,
	}

	blockchain.addBlock(newBlock)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Block added successfully"})
}

func main() {
	blockchain = Blockchain{
		Chain:      []Block{},
		Difficulty: 4,
	}

	genesisBlock := blockchain.createGenesisBlock()
	blockchain.Chain = append(blockchain.Chain, genesisBlock)

	http.HandleFunc("/chain", handleGetChain)
	http.HandleFunc("/addBlock", handleAddBlock)

	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
