package main

/**
Les types de données personnalisés sont définis à l'aide de la fonctionnalité structs de Go.
Ils sont définis à l'aide du mot-clé `type`, d'un nom et du mot-clé `struct`. Des accolades encadrent une série de champs,
chacun étant défini par un nom et un type. Les champs de même type peuvent être déclarés ensemble.
**/

type Product struct {
	name, category string
	price          float64
}

/**
Si un champ est défini sans nom, il est appelé champ incorporé et on y accède en utilisant le nom de son type.
Le type de structure StockLevel possède deux champs. Le premier champ est imbriqué et est défini uniquement à l'aide d'un type,
à savoir le type de structure Product.
**/

type StockLevel struct {
	Product
	count int
}

/**
Les noms de champs doivent être uniques pour chaque type de structure, ce qui signifie que nous ne pouvons définir qu'un seul champ
imbriqué pour un type spécifique. Si nous devons définir deux champs du même type, nous devrons en nommer un.
**/

type StockLevelAlternate struct {
	Product
	Alternate Product
	count     int
}

type Item struct {
	name     string
	category string
	price    float64
}
