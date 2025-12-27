package store

/**
Le type SpecialDeal définit un champ imbriqué *Product. Cette combinaison entraîne des doublons, car les deux types définissent des
champs Name et price. On trouve également une fonction constructeur et une méthode GetDetails, qui renvoie les valeurs des champs Name
et price, ainsi que le résultat de la méthode Price.
**/

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}

func (d *SpecialDeal) GetDetails() (string, float64, float64) {
	return d.Name, d.price, d.Price(0)
}

/**
Si l'on souhaite pouvoir appeler la méthode `Price` et obtenir un résultat qui dépende du champ `SpecialDeal.price`, l'on doit définir
une nouvelle méthode.
**/

func (d *SpecialDeal) Price(taxRate float64) float64 {
	return d.price
}
