package main

import "time"

// Account structure to hold account information
type Account struct {
	Name         string
	Balance      float64
	Transactions []Transaction
}

// Transaction structure to hold transaction information
type Transaction struct {
	Description string
	Amount      float64
	Date        time.Time
}
