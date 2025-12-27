package store

type Product struct {
	Name, Category string
	price          float64
}

func (p *Product) Price(taxRate float64) float64 {
	return p.price + p.price*taxRate
}

/**
Une convention courante consiste à définir une fonction constructeur dont le nom est New<Type> (ou new<Type>).
Les fonctions constructeurs ne sont qu'une convention, et leur utilisation n'est pas obligatoire, ce qui signifie que les
types exportés peuvent être créés en utilisant la syntaxe littérale, à condition qu'aucune valeur ne soit attribuée aux champs non exportés.

Il convient d'utiliser les constructeurs chaque fois qu'ils sont définis, car ils facilitent la gestion des modifications apportées
à la manière dont les valeurs sont créées et parce qu'ils garantissent que les champs sont correctement initialisés.
**/

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
