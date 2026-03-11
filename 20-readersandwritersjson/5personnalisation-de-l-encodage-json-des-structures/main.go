package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Personnalisation de l'encodage JSON des structures

L'encodage d'une structure peut être personnalisé à l'aide des étiquettes (tags) de structure, des chaînes de caractères littérales placées
après les champs. Ces étiquettes font partie de la prise en charge de la réflexion en Go. Les étiquettes suivent les champs et permettent
de modifier deux aspects de l'encodage JSON d'un champ.

Référence structure `DiscountedProduct`
L'étiquette struct suit un format spécifique (exemple ci-dessous).
Le terme « json » est suivi de deux points, puis du nom à utiliser lors de l’encodage du champ, le tout entre guillemets.
L'étiquette entière est encadrée par des accents graves.

Référence structure `OwnerProduct`
Omission d'un champ
Le processus d'encodage ignore les champs décorés d'une étiquette spécifiant un tiret (le caractère -) dans le nom.
La nouvelle étiquette indique à l'encodeur d'ignorer le champ Username lors de la création de la représentation JSON d'une valeur `OwnerProduct`.

Référence structure `LocationProduct`
Omission des champs non assignés
Par défaut, le processus d'encodage JSON inclut les champs de structure, même lorsqu'aucune valeur ne leur a été attribuée.
Pour omettre un champ nul, le mot-clé `omitempty` est ajouté à l'étiquette du champ. Ce mot-clé est séparé du nom du champ par une virgule, sans espace.

référence structure `PriceProduct`
Pour ignorer un champ nul sans modifier son nom ni sa promotion, spécifiez le mot-clé `omitempty` sans nom.
Le processus d'encodage affichera les champs `Product` si une valeur a été attribuée au champ intégré et omettra le champ si
aucune valeur n'a été attribuée.

Référence structure `TaxProduct`
Encodage forcé des champs sous forme de chaînes de caractères
Les étiquettes de `struct` permettent d'encoder de force la valeur d'un champ sous forme de chaîne de caractères,
en remplaçant l'encodage par défaut du type de champ.
L'ajout du mot-clé `string` remplace l'encodage par défaut et produit une chaîne pour le champ `Tax`.
**/

func main() {
	fmt.Println("#1 Encode d'une structure imbriquée avec tags...")
	var writer1 strings.Builder
	encoder1 := json.NewEncoder(&writer1)
	dp := DiscountedProduct{
		Product:  &kayak,
		Discount: 10.50,
	}
	encoder1.Encode(dp)
	fmt.Println(writer1.String())

	fmt.Println("#2 Encode d'une structure imbriquée avec tags ignorés...")
	var writer2 strings.Builder
	encoder2 := json.NewEncoder(&writer2)
	op := OwnerProduct{
		Product:  &kayak,
		Username: "willbrid",
	}
	encoder2.Encode(op)
	fmt.Println(writer2.String())

	fmt.Println("#3 Encode d'une structure imbriquée avec tags non ignorés pour les champs non assignés...")
	var writer3 strings.Builder
	encoder3 := json.NewEncoder(&writer3)
	dp1 := DiscountedProduct{Discount: 10.50}
	encoder3.Encode(dp1) // Le champ non assigné `product` est inclu
	fmt.Println(writer3.String())

	fmt.Println("#4 Encode d'une structure imbriquée avec tags ignorés pour les champs non assignés...")
	var writer4 strings.Builder
	encoder4 := json.NewEncoder(&writer4)
	lp := LocationProduct{Location: "Europe"}
	encoder4.Encode(lp)
	fmt.Println(writer4.String())

	fmt.Println("#5 Encode d'une structure imbriquée sans changer de nom avec tags ignorés pour les champs non assignés...")
	var writer5 strings.Builder
	encoder5 := json.NewEncoder(&writer5)
	pp1 := PriceProduct{Product: &kayak, Price: 700.00}
	encoder5.Encode(pp1)
	fmt.Println(writer5.String())

	fmt.Println("#6 Encode d'une structure imbriquée sans changer de nom avec tags ignorés pour les champs non assignés...")
	var writer6 strings.Builder
	encoder6 := json.NewEncoder(&writer6)
	pp2 := PriceProduct{Price: 700.00}
	encoder6.Encode(pp2)
	fmt.Println(writer6.String())

	fmt.Println("#7 Encode d'une structure imbriquée avec encodage forcé du type de champs...")
	var writer7 strings.Builder
	encoder7 := json.NewEncoder(&writer7)
	tp := TaxProduct{Product: &kayak, Tax: 15.00}
	encoder7.Encode(tp)
	fmt.Println(writer7.String())
}
