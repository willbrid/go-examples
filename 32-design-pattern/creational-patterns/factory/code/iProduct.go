package main

type IProduct interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}
