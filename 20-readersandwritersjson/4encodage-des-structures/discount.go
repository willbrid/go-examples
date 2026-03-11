package main

// Le type de structure `DiscountedProduct` définit un champ `Product` intégré
type DiscountedProduct struct {
	*Product
	Discount float64
}
