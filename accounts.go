package main

import "fmt"

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
