package main

import "fmt"

/**
Un modèle peut être utilisé pour rechercher des valeurs dans une chaîne contenant des caractères non obligatoires.

L'analyse par modèle est moins flexible que l'utilisation d'une expression régulière, car la chaîne analysée ne peut contenir que des
valeurs séparées par des espaces. Cependant, l'utilisation d'un modèle peut s'avérer utile si nous ne souhaitons extraire que certaines
valeurs d'une chaîne et que nous ne voulons pas définir de règles de correspondance complexes.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	var (
		name     string
		category string
		price    float64
	)

	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)

	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error :%v", err.Error())
	}
}
