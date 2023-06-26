package main

import (
	"fmt"
	"time"
)

type Token struct {
	Owner    string
	Quantity int
}

type Property struct {
	ID          int
	Address     string
	Description string
	Price       float64
	Tokens      []Token
}

type RealEstatePlatform struct {
	Properties map[int]Property
}

func NewRealEstatePlatform() *RealEstatePlatform {
	return &RealEstatePlatform{
		Properties: make(map[int]Property),
	}
}

func (p *RealEstatePlatform) CreateProperty(id int, address, description string, price float64) {
	property := Property{
		ID:          id,
		Address:     address,
		Description: description,
		Price:       price,
		Tokens:      make([]Token, 0),
	}

	p.Properties[id] = property
}

func (p *RealEstatePlatform) BuyTokens(propertyID int, buyer string, quantity int) error {
	property, ok := p.Properties[propertyID]
	if !ok {
		return fmt.Errorf("property with ID %d does not exist", propertyID)
	}

	if quantity <= 0 {
		return fmt.Errorf("quantity should be a positive number")
	}

	if quantity > len(property.Tokens) {
		return fmt.Errorf("insufficient tokens available for purchase")
	}

	property.Tokens = property.Tokens[:len(property.Tokens)-quantity]

	for i := 0; i < quantity; i++ {
		token := Token{
			Owner:    buyer,
			Quantity: 1,
		}
		property.Tokens = append(property.Tokens, token)
	}

	p.Properties[propertyID] = property
	return nil
}

func main() {
	platform := NewRealEstatePlatform()

	platform.CreateProperty(1, "123 Main St", "Beautiful house", 500000)
	platform.CreateProperty(2, "456 Elm St", "Spacious apartment", 300000)

	err := platform.BuyTokens(1, "John", 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = platform.BuyTokens(2, "Alice", 1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	property1 := platform.Properties[1]
	property2 := platform.Properties[2]

	fmt.Println("Property 1 tokens:", property1.Tokens)
	fmt.Println("Property 2 tokens:", property2.Tokens)
}
