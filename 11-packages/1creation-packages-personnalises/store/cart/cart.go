package cart

import "packages/store"

/**
Les packages peuvent être définis à l'intérieur d'autres packages, ce qui facilite la décomposition des fonctionnalités complexes en
autant d'unités que possible.

L'instruction `package` s'utilise comme pour n'importe quel autre package, sans qu'il soit nécessaire d'inclure le nom du package parent
ou englobant. En revanche, toute dépendance à des packages personnalisés doit inclure le chemin complet du package.
**/

type Cart struct {
	CustomerName string
	Products     []store.Product
}

func (c *Cart) GetTotal() (total float64) {
	for _, p := range c.Products {
		total += p.Price()
	}

	return
}
