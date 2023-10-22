package main

import (
	"fmt"
	"time"
)

func viewTransactions() {
	var index int
	fmt.Println("Enter the index of the account to view transactions for:")
	fmt.Scan(&index)
	if index >= 0 && index < len(accounts) {
		for i, transaction := range accounts[index].Transactions {
			fmt.Printf("%d. %s: $%.2f on %s\n", i, transaction.Description, transaction.Amount, transaction.Date.Format("2006-01-02"))
		}
	} else {
		fmt.Println("Invalid index.")
	}
}

func manageMoney() {
	var userInput string
	fmt.Println("Select a numeric option:")
	fmt.Println("1. Add money")
	fmt.Println("2. Remove money")
	fmt.Scanln(&userInput)

	switch userInput {
	case "1", "2":
		var index int
		var amount float64
		var description string
		fmt.Println("Enter the index of the account:")
		fmt.Scan(&index)
		if index >= 0 && index < len(accounts) {
			fmt.Println("Enter the amount:")
			fmt.Scan(&amount)
			fmt.Println("Enter a description for the transaction:")
			fmt.Scanln() // Consume leftover newline character
			fmt.Scanln(&description)
			date := time.Now() // gets the current date and time
			if userInput == "1" {
				accounts[index].Balance += amount
				accounts[index].Transactions = append(accounts[index].Transactions, Transaction{Description: description, Amount: amount, Date: date})
				fmt.Println("Money added.")
			} else {
				if accounts[index].Balance >= amount {
					accounts[index].Balance -= amount
					accounts[index].Transactions = append(accounts[index].Transactions, Transaction{Description: description, Amount: -amount, Date: date})
					fmt.Println("Money removed.")
				} else {
					fmt.Println("Insufficient balance.")
				}
			}
		} else {
			fmt.Println("Invalid index.")
		}
	default:
		fmt.Println("Invalid action")
	}
}
