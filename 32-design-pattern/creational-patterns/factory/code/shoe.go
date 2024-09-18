package main

type Shoe struct {
	Product
}

func newShoe() IProduct {
	return &Shoe{
		Product: Product{
			name:  "shoe product",
			price: 5,
		},
	}
}
