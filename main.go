package main

import (
	"fmt"
	"time"
	"os"
)

// Models
type Account struct {
	Name         string
	Balance      float64
	Transactions []Transaction
}

type Transaction struct {
	Description string
	Amount      float64
	Date        time.Time
}

// Globals
var accounts []Account

// Main
func main() {
	args := os.Args[1:]
	switch args[0] {
	case "-cli":
		cli()
	case "-web":
		web()
	default:
		fmt.Println("Invalid action")
	}
}

// Fascades
func cli() {
	var userInput string
	for {
		fmt.Println("Select a numeric option:")
		fmt.Println("1. Manage accounts")
		fmt.Println("2. Manage money")
		fmt.Println("3. Display accounts")
		fmt.Println("4. View transactions")
		fmt.Println("5. Exit")
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			manageAccounts()
		case "2":
			manageMoney()
		case "3":
			displayAccounts()
		case "4":
			viewTransactions()
		case "5":
			return // Exit the program
		default:
			fmt.Println("Invalid action")
		}
	}
}

func web() {
	fmt.Println("Not implemented yet")
}

// Functions
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

func manageAccounts() {
	var userInput string
	fmt.Println("Select a numeric option:")
	fmt.Println("1. Add an account")
	fmt.Println("2. Remove an account")
	fmt.Scanln(&userInput)

	switch userInput {
	case "1":
		var name string
		fmt.Println("Enter the name of the account:")
		fmt.Scanln(&name)
		accounts = append(accounts, Account{Name: name, Balance: 0})
		fmt.Println("Account added.")
	case "2":
		var index int
		fmt.Println("Enter the index of the account to remove:")
		fmt.Scan(&index)
		if index >= 0 && index < len(accounts) {
			accounts = append(accounts[:index], accounts[index+1:]...)
			fmt.Println("Account removed.")
		} else {
			fmt.Println("Invalid index.")
		}
	default:
		fmt.Println("Invalid action")
	}
}

func displayAccounts() {
	for i, account := range accounts {
		fmt.Printf("%d. %s: $%.2f\n", i, account.Name, account.Balance)
	}
}
