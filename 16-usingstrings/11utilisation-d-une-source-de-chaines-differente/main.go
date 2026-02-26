package main

import "fmt"

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	var (
		name     string
		category string
		price    float64
	)

	source := "Lifejacket Watersports 48.95"
	fmt.Print("Enter text to scan : ")
	n, err := fmt.Sscan(source, &name, &category, &price)

	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error :%v", err.Error())
	}
}
