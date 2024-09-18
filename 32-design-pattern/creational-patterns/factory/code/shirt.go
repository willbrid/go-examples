package main

type Shirt struct {
	Product
}

func newShirt() IProduct {
	return &Shirt{
		Product: Product{
			name:  "shirt product",
			price: 3,
		},
	}
}
