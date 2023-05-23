package main

import "sort"

type Product struct {
	Name  string
	Price float64
}

type ProductSlice []Product

type ProductSliceName struct{ ProductSlice }

type ProductComparison func(p1, p2 Product) bool

type ProductSliceFlex struct {
	ProductSlice
	ProductComparison
}

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

func ProductSlicesByName(p []Product) {
	sort.Sort(ProductSliceName{p})
}
func (p ProductSliceName) Less(i, j int) bool {
	return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

// Définition d'un paramètre pour la fonction de comparaison
func (flex ProductSliceFlex) Less(i, j int) bool {
	return flex.ProductComparison(flex.ProductSlice[i], flex.ProductSlice[j])
}

func SortWith(prods []Product, f ProductComparison) {
	sort.Sort(ProductSliceFlex{prods, f})
}
