package main

func printDetails(values ...Product) {
	for _, elem := range values {
		Printfln("Product: Name: %v, Category: %v, Price: %v", elem.Name, elem.Category, elem.Price)
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	printDetails(product)
}
