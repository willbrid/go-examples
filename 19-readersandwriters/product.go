package main

type Product struct {
	Name, Category string
	Price          float64
}

var kayak Product = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}
