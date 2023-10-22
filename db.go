package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Fatal(err)
	}

	createAccountTable := `CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		balance REAL
	);`
	_, err = db.Exec(createAccountTable)
	if err != nil {
		log.Fatal(err)
	}

	createTransactionTable := `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account_id INTEGER,
		description TEXT,
		amount REAL,
		date TEXT,
		FOREIGN KEY (account_id) REFERENCES accounts (id)
	);`
	_, err = db.Exec(createTransactionTable)
	if err != nil {
		log.Fatal(err)
	}
}
