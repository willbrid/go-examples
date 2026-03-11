package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Encodage des structures
Une valeur de type `Encoder` exprime les valeurs des structures sous forme d’objets JSON, en utilisant les noms des champs exportés comme clés et
leurs valeurs comme valeurs. Les champs non exportés sont ignorés.

Cet exemple ci-dessous encode la valeur de la structure `Product` nommée `Kayak`. La structure `Product` définit les champs exportés
`Name`, `Category` et `Price`.


Comprendre l'effet de la promotion dans l'encodage JSON
Lorsqu'une structure définit un champ imbriqué qui est également une structure, les champs de la structure imbriquée sont promus et
encodés comme s'ils étaient définis par le type englobant.
Dans le résultat de l'exemple #2 ci-dessous, la valeur de type `Encoder` encode et met en avant les champs `Product` dans la sortie JSON.
Notons que l'exemple #2 encode un pointeur vers la valeur de la structure (`Product`). La fonction `Encode` suit ce pointeur et encode
la valeur à son emplacement, ce qui signifie que le code de l'exemple #2 encode la valeur de `DiscountedProduct` sans en créer de copie.
**/

func main() {
	fmt.Println("#1 Encode d'une structure simple...")
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(kayak)
	fmt.Println(writer.String())

	fmt.Println("#2 Encode d'une structure imbriquée...")
	var writer1 strings.Builder
	encoder1 := json.NewEncoder(&writer1)
	dp := DiscountedProduct{
		Product:  &kayak,
		Discount: 10.50,
	}
	encoder1.Encode(dp)
	fmt.Println(writer1.String())
}
