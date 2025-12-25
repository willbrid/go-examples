package main

type Product struct {
	name, category string
	price          float64
}

/**
Pour implémenter une interface, toutes les méthodes spécifiées par l'interface doivent être définies pour un type struct.
Les méthodes implémentées doivent avoir le même nom, les mêmes types de paramètres et les mêmes types de résultats.
**/

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}
