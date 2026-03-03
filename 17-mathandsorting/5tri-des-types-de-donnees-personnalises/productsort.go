package main

import "sort"

/**
Le type `ProductSlice` est un alias pour une slice de `Product` et c'est le type pour lequel les méthodes de l'interface (Len, Less, Swap)
ont été implémentées. Outre ces méthodes, deux fonctions ont été définies :

- une fonction `ProductSlices` qui accepte une slice de `Product`, la convertit en type `ProductSlice` et la
transmet comme argument à la fonction `Sort` du package `sort`.

- une fonction `ProductSlicesAreSorted` qui appelle la fonction `IsSorted`.

Les noms de ces fonctions suivent la convention établie par le package `sort`, qui consiste à faire suivre le nom du type alias par la lettre `s`.
**/

type Product struct {
	Name  string
	Price float64
}

type ProductSlice []Product

// sort.Sort trie les données dans l'ordre croissant tel que déterminé par la méthode Less
func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}

func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}

func (products ProductSlice) Len() int {
	return len(products)
}

func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}

func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

/**
La composition de types peut être utilisée pour trier un même type de structure en utilisant différents champs.
Un type de structure est défini pour chaque champ de structure pour lequel un tri est requis, avec un champ `ProductSlice` intégré.

La fonctionnalité de composition de types implique que les méthodes définies pour le type `ProductSlice` sont étendues au type englobant.
Une nouvelle méthode `Less` est définie pour le type englobant ; elle servira à trier les données selon un champ différent.

La dernière étape consiste à définir une fonction qui effectuera une conversion d'une slice de produit vers le nouveau type et appellera
la fonction `Sort`.
**/

type ProductSliceName struct{ ProductSlice }

func ProductSlicesByName(p []Product) {
	sort.Sort(ProductSliceName{p})
}
func (p ProductSliceName) Less(i, j int) bool {
	return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

/**
Une autre approche consiste à spécifier l'expression utilisée pour comparer les éléments en dehors de la fonction de tri.

Un nouveau type, nommé ProductSliceFlex, est créé. Il combine les données et la fonction de comparaison, ce qui permet à cette approche de s'intégrer
à la structure des fonctions définies par le package `sort`. Une méthode `Less` est définie pour le type `ProductSliceFlex `; elle appelle la fonction
de comparaison. Enfin, la fonction `SortWith` combine les données et la fonction en une valeur ProductSliceFlex et la transmet à la fonction `sort.Sort`.
**/

type ProductComparison func(p1, p2 Product) bool

type ProductSliceFlex struct {
	ProductSlice
	ProductComparison
}

func (flex ProductSliceFlex) Less(i, j int) bool {
	return flex.ProductComparison(flex.ProductSlice[i], flex.ProductSlice[j])
}

func SortWith(prods []Product, f ProductComparison) {
	sort.Sort(ProductSliceFlex{prods, f})
}
