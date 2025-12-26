// Package store provides types and methods
// commonly required for online sales
package store

/**
Le package personnalisé est défini à l'aide du mot-clé `package` suivi de son nom (exemple `package store`).
Le nom spécifié par l'instruction `package` doit correspondre au nom du dossier dans lequel les fichiers de code sont créés
(dans l'exemple le dossier `store`)

Les commentaires doivent être simples et descriptifs, et la convention est de commencer le commentaire par le nom de la fonctionnalité.
Go prend également en charge un commentaire qui décrit un package entier et qui apparaît avant le mot-clé du `package`.

Go utilise une approche inhabituelle pour le contrôle d'accès. Go examine la première lettre du nom des fonctionnalités (types, fonctions
et méthodes) dans un fichier de code. Si cette première lettre est minuscule, la fonctionnalité est utilisable uniquement au sein du package
qui la définit. Les fonctionnalités sont exportées pour être utilisées en dehors du package lorsqu'elles commencent par une majuscule.

Le nom du type struct est `Product`, ce qui signifie qu'il est utilisable en dehors du package `store`. Les noms des champs `Name` et
`Category` commencent également par une majuscule, ce qui signifie qu'ils sont aussi exportés. Le champ `price` commence par une minuscule,
ce qui signifie qu'il est accessible uniquement au sein du package `store`.

Les méthodes suivent une convention de nommage typique pour les méthodes exportées qui accèdent à un champ non exporté, de sorte que
la méthode `Price` renvoie la valeur du champ et la méthode `SetPrice` attribue une nouvelle valeur.
**/

var standardTax *TaxRate = newTaxRate(0.25, 0.2)

// Product describes an item for sale
type Product struct {
	Name, Category string // Name and type of the product
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}

func (p *Product) PriceWithCategory() float64 {
	return standardTax.calcTaxWithCategory(p)
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
