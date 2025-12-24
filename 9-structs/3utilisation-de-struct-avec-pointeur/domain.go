package main

type Supplier struct {
	name, city string
}

type Product struct {
	name, category string
	price          float64
}

type Article struct {
	name, category string
	price          float64
	*Supplier
}

/**
Une fonction constructeur est chargée de créer des valeurs d'une structure à partir des valeurs reçues via les paramètres.
Les fonctions constructeurs permettent de créer des valeurs de structure de manière cohérente. Elles sont généralement nommées `new` ou `New`,
suivi du type de la structure ; ainsi, la fonction constructeur pour créer des valeurs de type `Product` s'appelle `newProduct`.

Les fonctions constructeurs renvoient des pointeurs vers la structure, et l'opérateur d'adresse est utilisé directement avec
la syntaxe littérale des structures.

L'avantage d'utiliser des fonctions constructeurs réside dans la cohérence : toute modification apportée au processus de construction est
ainsi répercutée sur toutes les valeurs de structure créées par la fonction.
**/

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

/**
un champ imbriqué a été ajouté au type `Article` qui utilise le type `Supplier`.
La fonction newArticle est créée et accepte un pointeur vers un type `Supplier`.
**/

func newArticle(name, category string, price float64, supplier *Supplier) *Article {
	return &Article{name, category, price, supplier}
}
