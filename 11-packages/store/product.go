/**
Go prend également en charge un commentaire décrivant un package complet, qui apparaît avant le mot-clé « package ».
Ces commentaires sont traités par l'outil "go doc", qui génère la documentation du code.
**/

// Package store provides types and methods
// commonly required for online sales
package store

var standardTax *TaxRate = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name, Category string
	price          float64
}

/**
Go a une approche inhabituelle du contrôle d'accès. Au lieu de s'appuyer sur des mots-clés dédiés, comme public et private,
Go examine la première lettre des noms donnés aux fonctionnalités dans un fichier de code, telles que les types, les fonctions et les méthodes.
Si la première lettre est en minuscule, la fonctionnalité ne peut être utilisée que dans le package qui la définit.
Les fonctionnalités sont exportées pour une utilisation en dehors du package en leur donnant une première lettre en majuscule.
**/

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) PriceStandardTax() float64 {
	return standardTax.calcTax(p)
}

func (p *Product) setPrice(newPrice float64) {
	p.price = newPrice
}
