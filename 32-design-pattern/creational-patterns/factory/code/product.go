package main

type Product struct {
	name  string
	price int
}

func (p *Product) setName(name string) {
	p.name = name
}

func (p *Product) setPrice(price int) {
	p.price = price
}

func (p *Product) getName() string {
	return p.name
}

func (p *Product) getPrice() int {
	return p.price
}
