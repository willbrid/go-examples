package main

type Product struct {
	name, category string
	price          float64
}

type StockLevel struct {
	Product
	Alternate Product
	count     int
}
