package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Investor struct {
	Name                string    `json:"name"`
	DateOfBirth         time.Time `json:"dateOfBirth"`
	Address             string    `json:"address"`
	IdentificationNumber string    `json:"identificationNumber"`
}

func main() {
	router := gin.Default()

	// Define an endpoint for investor identity verification
	router.POST("/verify", verifyInvestorIdentityHandler)

	router.Run(":8080")
}

func verifyInvestorIdentityHandler(c *gin.Context) {
	var investor Investor
	if err := c.ShouldBindJSON(&investor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := verifyInvestorIdentity(investor)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Identity verification successful!", "investor": investor})
}

func verifyInvestorIdentity(investor Investor) error {
	// Perform identity verification checks here
	if !isAgeEligible(investor.DateOfBirth) {
		return fmt.Errorf("investor must be at least 18 years old")
	}

	// Additional checks can be implemented, such as verifying identification documents, performing anti-money laundering (AML) checks, etc.

	return nil
}

func isAgeEligible(dateOfBirth time.Time) bool {
	minimumAge := 18
	now := time.Now().UTC()
	age := now.Year() - dateOfBirth.Year()
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	return age >= minimumAge
}
