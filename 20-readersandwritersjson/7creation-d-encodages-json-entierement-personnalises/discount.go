package main

import "encoding/json"

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"`
}

/**
L'on définit une map avec des clés de type chaîne de caractères et l'on utilise l'interface vide pour les valeurs.
Cela me permet de construire le JSON en ajoutant des paires clé-valeur à la map, puis de transmettre cette map à la fonction `json.Marshal`,
qui utilise le support intégré pour encoder chacune des valeurs qu'elle contient.
**/

func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]any{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}
