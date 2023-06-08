package main

import "encoding/json"

type DiscountedProduct struct {
	*Product
	Discount float64
}

/*
*
Le mot clé json est suivi de deux-points, suivi du nom qui doit être utilisé lorsque le champ est encodé, entre guillemets doubles.
La balise entière est entourée de backticks.
*
*/
type DiscountedProduct1 struct {
	*Product `json:"product"`
	Discount float64
}

// L'encodeur ignore les champs décorés d'une balise qui spécifie un trait d'union (le caractère -) pour le nom.
type DiscountedProduct2 struct {
	*Product `json:"product"`
	Discount float64 `json:"-"`
}

// Pour omettre un champ nul lors de l'encodage en JSON, le mot-clé omitempty est ajouté à la balise du champ
type DiscountedProduct3 struct {
	*Product `json:"product,omitempty"`
	Discount float64
}

// Pour ignorer un champ nul sans changer le nom ou la promotion du champ, spécifions le mot-clé omitempty sans nom comme suit
type DiscountedProduct4 struct {
	*Product `json:",omitempty"`
	Discount float64
}

/*
*
Les balises de classe peuvent être utilisées pour forcer l'encodage d'une valeur de champ sous forme de chaîne,
en remplaçant l'encodage normal du type de champ
*
*/
type DiscountedProduct5 struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"` // Le type string sera utilisé
}

/*
*
Création d'encodages JSON entièrement personnalisés

Cette fonction json.Marshal encode la valeur spécifiée au format JSON. Les résultats sont le contenu JSON exprimé dans une tranche d'octet et
une erreur, qui indique tout problème d'encodage.
*
*/
func (dp *DiscountedProduct5) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}
