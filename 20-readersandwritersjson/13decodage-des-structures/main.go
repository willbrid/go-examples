package main

import (
	"encoding/json"
	"io"
	"strings"
)

/**
La structure clé-valeur des objets JSON peut être décodée en valeurs de structure Go.

Le décodeur décode l'objet JSON et utilise les clés pour définir les valeurs des champs de la structure exportée.
La casse des champs et des clés JSON n'a pas besoin de correspondre ; le décodeur ignore toute clé JSON sans champ de structure associé
et tout champ de structure sans clé JSON associée. Les objets JSON de l'exemple ci-dessous présentent une casse différente et
contiennent un nombre de clés variable par rapport aux champs de la structure Product.


Interdiction des clés inutilisées
Par défaut, le processus de décodage ignore les clés JSON sans champ de structure correspondant. Ce comportement peut être modifié en
appelant la méthode `DisallowUnknownFields`, ce qui déclenche une erreur lorsqu'une telle clé est rencontrée.

Dans l'exemple #1 ci-dessous l'un des objets JSON contient une clé `inStock pour laquelle il n'existe aucun champ Product correspondant.
Normalement, cette clé serait ignorée, mais comme la méthode `DisallowUnknownFields` a été appelée, elle ne l'est pas.

Utilisation des étiquettes (tags) de structure pour contrôler le décodage
Les clés utilisées dans un objet JSON ne correspondent pas toujours aux champs définis par les structures d'un projet Go.
Dans ce cas, les étiquettes de structure permettent de faire correspondre les données JSON aux structures.
**/

func main() {
	reader := strings.NewReader(`
		{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports" }
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
	`)
	decoder := json.NewDecoder(reader)
	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("#0 Error: %v", err.Error())
			}
			break
		} else {
			Printfln("#0 Name: %v, Category: %v, Price: %v", val.Name, val.Category, val.Price)
		}
	}

	reader1 := strings.NewReader(`
		{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports" }
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
	`)
	decoder1 := json.NewDecoder(reader1)
	decoder1.DisallowUnknownFields()
	for {
		var val1 Product
		err := decoder1.Decode(&val1)
		if err != nil {
			if err != io.EOF {
				Printfln("#1 Error: %v", err.Error())
			}
			break
		} else {
			Printfln("#1 Name: %v, Category: %v, Price: %v", val1.Name, val1.Category, val1.Price)
		}
	}

	reader2 := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	decoder2 := json.NewDecoder(reader2)
	for {
		var val2 DiscountedProduct
		err := decoder2.Decode(&val2)
		if err != nil {
			if err != io.EOF {
				Printfln("#2 Error: %v", err.Error())
			}
			break
		} else {
			Printfln("#2 Name: %v, Category: %v, Price: %v, Discount: %v", val2.Name, val2.Category, val2.Price, val2.Discount)
		}
	}
}
