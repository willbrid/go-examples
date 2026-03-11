package main

/**
Ce fichier définit une interface simple et une structure qui l'implémente, ainsi qu'une méthode pour la
structure `DiscountedProduct` afin qu'elle implémente l'interface.
**/

type Named interface {
	GetName() string
}

type Person struct {
	PersonName string
}

func (p *Person) GetName() string {
	return p.PersonName
}

func (p *DiscountedProduct) GetName() string {
	return p.Name
}
