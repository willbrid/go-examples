package main

type Article struct {
	name, category string
	price          float64
}

func (a *Article) getName() string {
	return a.name
}

func (a *Article) getCost(_ bool) float64 {
	return a.price
}
