package main

import "fmt"

/**
Le package fmt fournit des fonctions pour composer et écrire des chaînes de caractères.
Les fonctions de base :

`Print(...vals)` : Cette fonction accepte un nombre variable d'arguments et affiche leurs valeurs sur la sortie standard.
Des espaces sont ajoutés entre les valeurs qui ne sont pas des chaînes de caractères.

`Println(...vals)` : Cette fonction accepte un nombre variable d'arguments et affiche leurs valeurs sur la sortie standard,
séparées par des espaces et suivies d'un saut de ligne.

`Fprint(writer, ...vals)` : Cette fonction écrit un nombre variable d'arguments dans le flux d'écriture spécifié.
Des espaces sont ajoutés entre les valeurs qui ne sont pas des chaînes de caractères.

`Fprintln(writer, ...vals)` : Cette fonction écrit un nombre variable d'arguments dans le flux d'écriture spécifié,
suivi d'un saut de ligne. Des espaces sont ajoutés entre toutes les valeurs.
**/

func main() {
	fmt.Println("Product :", kayak.Name, "- Price :", kayak.Price)
	fmt.Print("Product :", kayak.Name, "- Price :", kayak.Price, "\n")
}
