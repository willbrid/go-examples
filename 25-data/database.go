package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}

func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if err == nil {
		Printfln("Opened database")
	}
	return
}

func InsertNewCategoryPrepare(db *sql.DB) (query *sql.Stmt) {
	query, _ = db.Prepare("INSERT INTO Categories (Name) VALUES (?)")
	return query
}

func ChangeProductCategoryPrepare(db *sql.DB) (query *sql.Stmt) {
	query, _ = db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")
	return query
}
