package main

type DiscountedProduct struct {
	*Product `json:"product"`
	Discount float64
}

type OwnerProduct struct {
	*Product `json:"product"`
	Username string `json:"-"`
}

type LocationProduct struct {
	*Product `json:"product,omitempty"`
	Location string `json:"location"`
}

type PriceProduct struct {
	*Product `json:",omitempty"`
	Price    float64 `json:"price"`
}

type TaxProduct struct {
	*Product `json:",omitempty"`
	Tax      float64 `json:",string"`
}
