package main

import (
	"encoding/json"
	"strconv"
)

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
Les clés utilisées dans un objet JSON ne correspondent pas toujours aux champs définis par les classes dans un projet Go.
Lorsque cela se produit, les balises struct peuvent être utilisées pour mapper entre les données JSON et la classe
*
*/
type DiscountedProduct6 struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"offer,string"` // mappage entre les données JSON et la classe et le type string sera utilisé
}

type DiscountedProduct7 struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:"offer,string"` // mappage entre les données JSON et la classe et le type string sera utilisé
}

/*
*
Création d'encodage JSON entièrement personnalisé

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

/*
*
Création de décodage JSON entièrement personnalisé

Unmarshal(data []byte, v)
Unmarshal analyse les données encodées en JSON et stocke le résultat dans la valeur pointée par v. Si v est nil ou n'est pas un pointeur,
Unmarshal renvoie une InvalidUnmarshalError.
Unmarshal utilise l'inverse des encodages utilisés par Marshal, allouant des maps, des tranches et des pointeurs selon les besoins,
avec les règles supplémentaires suivantes : pour démarshaler JSON dans un pointeur, Unmarshal gère d'abord le cas où le JSON est le littéral JSON nul.
Dans ce cas, Unmarshal définit le pointeur sur nil. Sinon, Unmarshal démarshale le JSON dans la valeur pointée par le pointeur. Si le pointeur est nil,
Unmarshal lui alloue une nouvelle valeur vers laquelle pointer.
*
*/
func (dp *DiscountedProduct7) UnmarshalJSON(data []byte) (err error) {
	mdata := map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)

	if dp.Product == nil {
		dp.Product = &Product{}
	}

	if err == nil {
		if name, ok := mdata["Name"].(string); ok {
			dp.Name = name
		}
		if category, ok := mdata["Category"].(string); ok {
			dp.Category = category
		}
		if price, ok := mdata["Price"].(float64); ok {
			dp.Price = price
		}
		if discount, ok := mdata["Offer"].(string); ok {
			fpval, fperr := strconv.ParseFloat(discount, 64)
			if fperr == nil {
				dp.Discount = fpval
			}
		}
	}
	return
}
