package main

import "encoding/json"

/**
L'étiquette appliquée au champ `Discount` indique au processus de décodage que la valeur de ce champ doit être obtenue à partir de la clé JSON
nommée `offer` et que la valeur sera analysée à partir d'une chaîne de caractères, au lieu du nombre JSON attendu pour une valeur Go float64.
**/

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"offer,string"`
}

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
