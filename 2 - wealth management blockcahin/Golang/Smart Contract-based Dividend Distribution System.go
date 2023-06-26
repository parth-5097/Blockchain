package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DividendDistribution represents a smart contract for dividend distribution
type DividendDistribution struct {
	TotalSupply       *big.Int                // Total supply of tokens
	BalanceOf         map[common.Address]*big.Int // Balance of tokens for each shareholder
	DividendPerShare  *big.Int                // Dividend amount per share
	Shareholders      []common.Address        // List of shareholders
	ShareholderShares map[common.Address]*big.Int // Number of shares held by each shareholder
}

// NewDividendDistribution initializes a new DividendDistribution contract
func NewDividendDistribution(totalSupply *big.Int) *DividendDistribution {
	return &DividendDistribution{
		TotalSupply:       totalSupply,
		BalanceOf:         make(map[common.Address]*big.Int),
		DividendPerShare:  big.NewInt(0),
		Shareholders:      []common.Address{},
		ShareholderShares: make(map[common.Address]*big.Int),
	}
}

// AddShareholder adds a new shareholder with the specified address and number of shares
func (dd *DividendDistribution) AddShareholder(address common.Address, shares *big.Int) {
	dd.Shareholders = append(dd.Shareholders, address)
	dd.ShareholderShares[address] = shares
}

// DistributeDividend distributes dividends to all shareholders
func (dd *DividendDistribution) DistributeDividend(dividendAmount *big.Int) {
	// Calculate dividend per share
	dividendPerShare := new(big.Int).Div(dividendAmount, dd.TotalSupply)

	// Update dividend per share in the contract
	dd.DividendPerShare = dividendPerShare

	// Distribute dividend to each shareholder
	for _, shareholder := range dd.Shareholders {
		shares := dd.ShareholderShares[shareholder]
		dividend := new(big.Int).Mul(dividendPerShare, shares)
		dd.BalanceOf[shareholder] = dividend
	}
}

// GetDividendBalance returns the dividend balance for a shareholder
func (dd *DividendDistribution) GetDividendBalance(address common.Address) *big.Int {
	return dd.BalanceOf[address]
}

// Example usage
func main() {
	// Create a new dividend distribution contract with total supply of 1000 tokens
	contract := NewDividendDistribution(big.NewInt(1000))

	// Add shareholders and their respective shares
	contract.AddShareholder(common.HexToAddress("0x0123456789abcdef0123456789abcdef0123456"), big.NewInt(100))
	contract.AddShareholder(common.HexToAddress("0xabcdef0123456789abcdef0123456789abcdef01"), big.NewInt(200))
	contract.AddShareholder(common.HexToAddress("0x456789abcdef0123456789abcdef0123456789ab"), big.NewInt(300))

	// Distribute a dividend of 1000 wei
	contract.DistributeDividend(big.NewInt(1000))

	// Get dividend balance for a shareholder
	balance := contract.GetDividendBalance(common.HexToAddress("0x0123456789abcdef0123456789abcdef0123456"))
	fmt.Println("Dividend balance:", balance)
}
