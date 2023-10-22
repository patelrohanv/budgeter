package main

import (
	"fmt"
)

// Global slice to hold accounts
var accounts []Account

func main() {
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
