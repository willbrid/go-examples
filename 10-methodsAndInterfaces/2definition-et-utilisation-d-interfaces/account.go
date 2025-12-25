package main

/**
Les types d'interface peuvent être utilisés pour les champs de structure, ce qui signifie que les champs peuvent se voir attribuer
des valeurs de n'importe quel type implémentant les méthodes définies par l'interface.

La structure Account possède un champ dépenses dont le type est une tranche de valeurs Expense, qui peut être utilisé comme
n'importe quel autre champ.
**/

type Account struct {
	accountNumber int
	expenses      []Expense
}
