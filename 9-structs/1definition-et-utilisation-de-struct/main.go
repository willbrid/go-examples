package main

import "fmt"

func main() {
	/**
	L'on peut créer une valeur à l'aide du type personnalisé, ce qui se fait en utilisant le nom du type de structure,
	suivi d'accolades contenant les valeurs des champs de structure (Product).
	Go est très pointilleux sur la syntaxe et générera une erreur si la valeur du dernier champ n'est pas suivie d'une
	virgule ou d'une accolade fermante.
	Les champs d'une structure sont accessibles via le nom donné à la variable. Ainsi, la valeur du champ « name » de la
	structure Product assignée à la variable « kayak » est accessible via `kayak.name`. De nouvelles valeurs peuvent être attribuées à
	un champ de structure en utilisant la même syntaxe.
	**/
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	fmt.Println("kayak :", kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Kayak changed price :", kayak.price)

	/**
	Il n'est pas nécessaire de fournir des valeurs pour tous les champs lors de la création d'une valeur de structure.
	Lorsqu'aucun champ n'est fourni, la valeur zéro est utilisée pour le type du champ.
	**/
	kayak1 := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println("kayak1 :", kayak1.name, kayak1.category, kayak1.price)
	kayak1.price = 300
	fmt.Println("Kayak1 changed price :", kayak1.price)

	/**
	Certains codes utilisent la fonction intégrée `new` pour créer des valeurs de structure. Le résultat est un pointeur vers une valeur de
	structure dont les champs sont initialisés à la valeur zéro de leur type.
	Et new(Product) est équivalent à `&Product{}`
	**/
	kayak2 := new(Product)
	fmt.Println("kayak2 :", kayak2.name, kayak2.category, kayak2.price)
	kayak2.price = 300
	fmt.Println("Kayak2 changed price :", kayak2.price)

	/**
	Les valeurs d'une structure peuvent être définies sans utiliser de noms, à condition que leurs types correspondent à l'ordre
	dans lequel les champs sont définis par le type de la structure.
	**/
	kayak3 := Product{"Kayak", "Watersports", 275.00}
	fmt.Println("kayak3 Name :", kayak3.name, "category :", kayak3.category, "price :", kayak3.price)

	/**
	Les champs intégrés sont accessibles via le nom du type de champ.
	**/
	stockItem := StockLevel{
		Product: Product{"kayak", "Watersports", 275.00},
		count:   100,
	}
	fmt.Println("StockItem name :", stockItem.Product.name)
	fmt.Println("StockItem count :", stockItem.count)

	stockItemAlternate := StockLevelAlternate{
		Product:   Product{"kayak", "Watersports", 275.00},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("stockItemAlternate name :", stockItemAlternate.Product.name)
	fmt.Println("stockItemAlternate alt name:", stockItemAlternate.Alternate.name)
	fmt.Println("stockItemAlternate count :", stockItemAlternate.count)

	/**
	Les valeurs de structure sont comparables si tous leurs champs peuvent être comparés.
	Les structures p1 et p2 ont la même valeur car tous leurs champs sont identiques. Les structures p1 et p3 ont la même valeur
	car les valeurs attribuées à leurs champs de catégorie sont différentes.
	Il est impossible de comparer des structures si leur type définit des champs de types incomparables.
	**/
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2 :", p1 == p2)
	fmt.Println("p1 == p3 :", p1 == p3)

	/**
	Un type struct peut être converti en n'importe quel autre type struct possédant les mêmes champs, c'est-à-dire que tous les champs
	ont le même nom et le même type et sont définis dans le même ordre et les champs de types incomparables.
	**/
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item :", prod == Product(item))

	prod1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item1 := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod1)
	writeName(item1)
}
