package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
*
Définir une structure anonyme et lui attribuer une valeur en une seule étape est une bonne pratique

Fonction avec pour type d'un paramètre une structure anonyme
Les types de structures anonymes sont définis sans utiliser de nom.
*
*/
func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name : ", val.name)
}

/**
Go ne permet pas d'utiliser des structures avec le mot-clé const, et le compilateur signalera une erreur si nous essayons de définir
une structure constante.
**/

/** Définir et utiliser une structure **/
type Product struct {
	name, category string
	price          float64
}

/*
*
Les noms de champ doivent être uniques avec le type struct, ce qui signifie que nous ne pouvons définir qu'un seul
champ imbriqué pour un type spécifique. Si nous avons besoin de définir deux champs du même type, alors nous devrons attribuer
un nom à l'un d'entre eux. Exemple avec **Alternate**
*
*/
type StockLevel struct { // Définition d'une structure imbriquant une autre structure
	Product   // Ce champ n'a pas de nom et fait référence à la structure Product
	Alternate Product
	count     int
}

/** Comprendre les structures et les pointeurs **/

/*
* Go peut suivre les pointeurs vers les champs struct sans avoir besoin d'un astérisque pour les variables
injectées en argument de fonction : plus besoin d'utiliser dans le corps de la fonction l'astérique pour
accéder à la valeur de cette variable pointeur sur une autre variable.
Cette fonctionnalité ne modifie pas le type de données du paramètre de fonction, qui est toujours : *Product,
*
*/
func calcTax(product *Product) {
	if product.price > 100 { // Au lieu (*product).price > 100
		product.price += product.price * 0.2 // Au lieu (*product).price += (*product).price * 0.2
	}
}

func calcTaxWithReturn(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}

	return product
}

/*
* Une fonction constructeur est responsable de la création de valeurs de structure à l'aide de valeurs reçues via des paramètres
Les fonctions constructeur sont utilisées pour créer des valeurs de structure de manière cohérente. Les fonctions constructeur
sont généralement nommées new ou New suivies du type de structure. Ainsi la fonction constructeur pour la création de valeurs Product
soit nommée newProduct.
*
*/
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

type Article struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newArticle(name, category string, price float64, supplier *Supplier) *Article {
	return &Article{name, category, price - 10, supplier}
}

func copyArticle(article *Article) Article {
	a := *article
	s := *article.Supplier
	a.Supplier = &s

	return a
}

func main() {
	fmt.Println("Hello, Structs")

	/** Définir et utiliser une structure **/
	fmt.Println("Définir et utiliser une structure")

	var kayak Product = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	fmt.Println("Produit : ", kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price : ", kayak.price)

	/**
	Aucune valeur initiale n'est fournie pour le champ price de la structure affectée à la variable kayak.
	Lorsqu'aucun champ n'est fourni, la valeur zéro du type de champ est utilisée. Dans l'exemple type zéro du champ price est 0,
	car le type du champ est float64.
	**/
	var kayak1 Product = Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println("Produit : ", kayak1.name, kayak1.category, kayak1.price)
	kayak1.price = 300
	fmt.Println("Changed price : ", kayak1.price)

	/**
	Les types zéro sont attribués à tous les champs si nous définissons une variable de type structure mais ne lui attribuons pas de valeur.
	**/
	var lifejacket Product
	fmt.Println("Name is zero value : ", lifejacket.name == "")
	fmt.Println("Category is zero value : ", lifejacket.category == "")
	fmt.Println("Price is zero value : ", lifejacket.price == 0)

	var kayak2 Product = Product{ // On se base sur la position des paramètres sans mentionner les noms des champs pour initialiser notre struct Product
		"Kayak",
		"Watersports",
		275.00,
	}
	fmt.Println("Name : ", kayak2.name)
	fmt.Println("Category : ", kayak2.category)
	fmt.Println("Price : ", kayak2.price)

	/**
	Si un champ est défini sans nom, il est appelé champ imbriqué et on y accède en utilisant le nom de son type.
	Les noms de champ doivent être uniques avec le type de structure, ce qui signifie que nous ne pouvons définir qu'un seul champ imbriqué
	pour un type spécifique. Si nous devons définir deux champs du même type, nous devrions attribuer un nom à l'un d'eux.
	**/
	var stockItem StockLevel = StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00}, // Les champs imbriqués sont accessibles en utilisant le nom du type de champ.
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("Name: ", stockItem.Product.name) // Les champs imbriqués sont accessibles en utilisant le nom du type de champ
	fmt.Println("Count : ", stockItem.count)
	fmt.Println("Alt Name : ", stockItem.Alternate.name)

	/** La comparaison est faite sur toutes les valeurs de champ.
	    Les valeurs de structure sont comparables si tous leurs champs sont comparables.
		Les structures ne peuvent pas être comparées si le type de structure définit des champs avec des types incomparables,
		tels que des tranches
		type Product struct {
	        name, category string
	        price float64
	        otherNames []string
	    }
		**/
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2 : ", p1 == p2)
	fmt.Println("p1 == p3 : ", p1 == p3)

	/**
	Un type de structure peut être converti en n'importe quel autre type de structure qui possède les mêmes champs,
	ce qui signifie que tous les champs ont le même nom et le même type et sont définis dans le même ordre
	**/
	type Item struct {
		name     string
		category string
		price    float64
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item : ", prod == Product(item)) // On convertit un Item en Product

	prod1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item1 := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod1)
	writeName(item1)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod1.name,
		ProductPrice: prod1.price,
	})
	fmt.Println(builder.String())

	var array1 [1]StockLevel = [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array : ", array1[0].Product.name)

	var slice1 []StockLevel = []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Slice : ", slice1[0].Product.name)

	var kvp1 map[string]StockLevel = map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Map : ", kvp1["kayak"].Product.name)

	/** Comprendre les structures et les pointeurs **/
	fmt.Println("Comprendre les structures et les pointeurs")
	var p4 Product = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	var p5 *Product = &p4
	p4.name = "Original Kayak"
	fmt.Println("P4 : ", p4.name)
	fmt.Println("P5 : ", (*p5).name)
	fmt.Println("P5 : ", p5.name) // Avec les pointeurs sur les structures, on peut omettre l'astérisque pour accéder à un champ de la structure

	calcTax(&p4)
	fmt.Println("Name : ", p4.name, " - Category : ", p4.category, " - Price : ", p4.price)

	// L'opérateur d'adresse est utilisé avant le type de structure
	var p6 *Product = &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax(p6)
	fmt.Println("Name : ", p6.name, " - Category : ", p6.category, " - Price : ", p6.price)

	var p7 *Product = calcTaxWithReturn(&Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})
	fmt.Println("Name : ", p7.name, " - Category : ", p7.category, " - Price : ", p7.price)

	// Une fonction constructeur est responsable de la création de valeurs de structure à l'aide de valeurs reçues via des paramètres
	// définition d'un tableaux au pointeur de Product
	var products [2]*Product = [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name : ", p.name, " - Category : ", p.category, " - Price : ", p.price)
	}

	var acme *Supplier = &Supplier{"Acme Co", "New York"}
	var articles [2]*Article = [2]*Article{
		newArticle("Kayak", "Watersports", 275, acme),
		newArticle("Hat", "Skiing", 42.50, acme),
	}
	for _, a := range articles {
		fmt.Println("Name : ", a.name, " - Category : ", a.category, " - Price : ", a.price, " - Supplier : ", a.Supplier.name, a.Supplier.city)
	}

	var acme1 *Supplier = &Supplier{"Acme Co", "New York"}           // On crée un pointeur sur la structure Supplier
	var p8 *Article = newArticle("Kayak", "Watersports", 275, acme1) // On crée un pointeur sur la structure Article
	var p9 Article = *p8                                             // On assigne la valeur de la variable pointée par le pointeur sur la structure Artcile p8 : il s'agit d'un objet de la structure Article
	p8.name = "Original Kayak"
	p8.Supplier.name = "BoatCo"
	for _, p := range []Article{*p8, p9} { // On crée un tableau slice d'article contenant deux articles : *p8 valeur de la variable pointée par le pointeur sur la structure Artcile et l'objet Article p9
		fmt.Println("Name : ", p.name, " - Supplier : ", p.Supplier.name, p.Supplier.city)
	}

	var acme2 *Supplier = &Supplier{"Acme Co", "New York"}
	var p10 *Article = newArticle("Kayak", "Watersports", 275, acme2)
	var p11 Article = copyArticle(p10) // Copie entière de l'objet Article
	p10.name = "Original Kayak"
	p10.Supplier.name = "BoatCo"
	for _, p := range []Article{*p10, p11} {
		fmt.Println("Name : ", p.name, " - Supplier : ", p.Supplier.name, p.Supplier.city) // Ici les objets Supplier sont différents
	}

	// Afficher la valeur zéro des variables de type Product
	var prod2 Product
	var prod2Ptr *Product
	fmt.Println("Value Product : ", prod2.name, prod2.category, prod2.price)
	fmt.Println("Pointer Product : ", prod2Ptr)

	var art Article = Article{Supplier: &Supplier{}}
	var artPtr *Article
	fmt.Println("Value Article : ", art.name, art.category, art.price, art.Supplier.name)
	fmt.Println("Pointer Article : ", artPtr)
}
