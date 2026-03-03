package main

/**
Pour trier les types de données personnalisés, le package `sort` définit une interface nommée de manière confuse `Interface`, qui spécifie les méthodes
ci-dessous :

Len() : Cette méthode renvoie le nombre d'éléments à trier.

Less(i, j) : Cette méthode renvoie vrai si l'élément à l'indice i doit apparaître avant l'élément à l'indice j dans la séquence triée.
Si Less(i, j) et Less(j, i) renvoient tous deux faux, les éléments sont considérés comme égaux.

Swap(i, j) : Cette méthode échange les éléments aux indices spécifiés.

NB: Les méthodes définies ci-dessus sont appliquées à la collection d'éléments de données à trier, ce qui implique l'introduction d'un alias de type
et de fonctions qui effectuent des conversions pour appeler les fonctions définies ci-dessous.


Lorsqu'un type définit les méthodes ci-dessus, il peut être trié à l'aide des fonctions ci-dessous définies par le package `sort` :

Sort(data) : Cette fonction utilise les méthodes ci-dessus pour trier les données spécifiées.

Stable(data) : Cette fonction utilise les méthodes ci-dessus pour trier les données spécifiées sans modifier l’ordre des éléments de même valeur.

IsSorted(data) : Cette fonction renvoie vrai si les données sont triées.

Reverse(data) : Cette fonction inverse l’ordre des données.
**/

func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(products)
	for _, p := range products {
		Printfln("products - Name: %v, Price: %.2f", p.Name, p.Price)
	}

	product1s := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlicesByName(product1s)
	for _, p := range product1s {
		Printfln("product1s - Name: %v, Price: %.2f", p.Name, p.Price)
	}

	// Les données sont triées en comparant le champ `Name`
	product2s := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	SortWith(product2s, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range product2s {
		Printfln("product2s - Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
