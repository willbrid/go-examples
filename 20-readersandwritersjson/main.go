package main

import (
	"encoding/json"
	"strings"
)

func main() {
	Printfln("Reading and Writing JSON Data")
	// Le package encoding/json prend en charge l'encodage et le décodage des données JSON
	/**
		bool : Les valeurs bool Go sont exprimées sous la forme JSON true ou false.
		string : Les valeurs de string Go sont exprimées sous forme de string JSON. Par défaut, les caractères HTML non sécurisés sont échappés.
		float32, float64 : Les valeurs à virgule flottante Go sont exprimées sous forme de nombres JSON.
		int, int<size> : Les valeurs entières Go sont exprimées sous forme de nombres JSON.
		uint, uint<size> : Les valeurs entières Go sont exprimées sous forme de nombres JSON.
		byte : Les octets (bytes) Go sont exprimés sous forme de nombres JSON.
		rune : Les valeurs de rune Go sont exprimées sous forme de nombres JSON.
		nil : La valeur Go nil est exprimée sous la forme de la valeur null JSON.
	 	Pointers : L'encodeur JSON suit les pointeurs et encode la valeur à l'emplacement du pointeur.
		**/
	var (
		b       bool    = true
		str     string  = "Hello"
		fval    float64 = 99.99
		ival    int     = 200
		pointer *int    = &ival
	)
	var writer1 strings.Builder
	// Cette fonction json.NewEncoder renvoie un Encoder, qui peut être utilisé pour encoder des données JSON et les écrire dans le Writer spécifié.
	var encoder1 *json.Encoder = json.NewEncoder(&writer1)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		// Cette méthode encoder.Encode encode la valeur spécifiée au format JSON et l'écrit dans le Writer.
		encoder1.Encode(val)
	}
	Printfln("Write 1 : %v", writer1.String())

	/**
	Les tranches (slice) et les tableaux Go sont encodés en tant que tableaux JSON, à l'exception que les tranches d'octets sont exprimées
	en chaînes encodées en base64. Les tableaux d'octets, cependant, sont codés comme un tableau de nombres JSON.
	**/
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	Printfln("[]byte(names[0]) : %v", []byte(names[0]))
	Printfln("byteArray : %v", byteArray)
	byteSlice := []byte(names[0])
	var writer2 strings.Builder
	var encoder2 *json.Encoder = json.NewEncoder(&writer2)
	encoder2.Encode(names)
	encoder2.Encode(numbers)
	encoder2.Encode(byteArray)
	encoder2.Encode(byteSlice)
	Printfln("Write 2 : %v", writer2.String())

	/**
	Les map Go sont encodées sous forme d'objets JSON, les clés de map étant utilisées comme clés d'objet.
	Les valeurs contenues dans la map sont encodées en fonction de leur type.
	**/
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	var writer3 strings.Builder
	var encoder3 *json.Encoder = json.NewEncoder(&writer3)
	encoder3.Encode(m)
	Printfln("Write 3 : %v", writer3.String())

	/**
	L'encodeur exprime les valeurs de classe sous forme d'objets JSON, en utilisant les noms de champ de structure
	exportés comme clés de l'objet et les valeurs de champ comme valeurs de l'objet.
	**/
	var writer4 strings.Builder
	var encoder4 *json.Encoder = json.NewEncoder(&writer4)
	encoder4.Encode(kayak)
	Printfln("Write 4 : %v", writer4.String())

	/**
	Lorsqu'une classe définit un champ incorporé qui est également une classe, les champs de la classe incorporée sont promus et encodés
	comme s'ils étaient définis par le type englobant.

	Notons que l'instruction encoder5.Encode encode un pointeur vers la valeur de la structure. La fonction Encode suit le pointeur et encode
	la valeur à son emplacement, ce qui signifie que l'instruction encoder5.Encode encode la valeur DiscountedProduct sans créer de copie.
	**/
	var writer5 strings.Builder
	var encoder5 *json.Encoder = json.NewEncoder(&writer5)
	dp := DiscountedProduct{
		Product:  &kayak,
		Discount: 10.50,
	}
	// L'encodeur promeut les champs Product dans la sortie JSON.
	encoder5.Encode(&dp)
	Printfln("Write 5 : %v", writer5.String())

	var writer6 strings.Builder
	var encoder6 *json.Encoder = json.NewEncoder(&writer6)
	dp1 := DiscountedProduct1{
		Product:  &kayak,
		Discount: 10.50,
	}
	// L'encodeur promeut les champs Product dans la sortie JSON.
	encoder6.Encode(&dp1)
	Printfln("Write 6 : %v", writer6.String())

	/**
	L'encodeur ignore les champs décorés d'une balise qui spécifie un trait d'union (le caractère -) pour le nom
	La nouvelle balise indique à l'encodeur d'ignorer le champ Discount lors de la création de la représentation JSON d'une valeur DIscountedProduct.
	**/
	var writer7 strings.Builder
	var encoder7 *json.Encoder = json.NewEncoder(&writer7)
	dp2 := DiscountedProduct2{
		Product:  &kayak,
		Discount: 10.50,
	}
	encoder7.Encode(&dp2)
	Printfln("Write 7 : %v", writer7.String())

	/**
	Par défaut, l'encodeur JSON inclut des champs de classe, même lorsqu'ils n'ont pas reçu de valeur.
	**/
	var writer8 strings.Builder
	var encoder8 *json.Encoder = json.NewEncoder(&writer8)
	dp3 := DiscountedProduct1{
		Discount: 10.50,
	}
	encoder8.Encode(&dp3)
	Printfln("Write 8 : %v", writer8.String())

	// Pour omettre un champ nul, le mot-clé omitempty est ajouté à la balise du champ
	var writer9 strings.Builder
	var encoder9 *json.Encoder = json.NewEncoder(&writer9)
	dp4 := DiscountedProduct3{
		Discount: 10.50,
	}
	encoder9.Encode(&dp4)
	Printfln("Write 9 : %v", writer9.String())

	// Pour ignorer un champ nul sans changer le nom ou la promotion du champ, spécifiez le mot-clé omitempty sans nom
	var writer10 strings.Builder
	var encoder10 *json.Encoder = json.NewEncoder(&writer10)
	dp5 := DiscountedProduct4{
		Discount: 10.50,
	}
	encoder10.Encode(&dp5)
	Printfln("Write 10 : %v", writer10.String())

	/**
	Forcer les champs à être encodés en tant que chaînes
	L'ajout du mot clé string remplace l'encodage par défaut et produit une chaîne pour le champ Discount
	**/
	var writer11 strings.Builder
	var encoder11 *json.Encoder = json.NewEncoder(&writer11)
	dp6 := DiscountedProduct5{
		Discount: 10.50,
	}
	encoder11.Encode(&dp6)
	Printfln("Write 11 : %v", writer11.String())

	/**
	L'encodeur JSON peut être utilisé sur des valeurs affectées à des variables d'interface, mais c'est le type dynamique qui est encodé
	La tranche de valeurs nommées contient différents types dynamiques.
	**/
	var writer12 strings.Builder
	var encoder12 *json.Encoder = json.NewEncoder(&writer12)
	dp7 := DiscountedProduct1{
		Product:  &kayak,
		Discount: 10.50,
	}
	namedItems := []Named{&dp7, &Person{PersonName: "Alice"}}
	encoder12.Encode(namedItems)
	Printfln("Write 12 : %v", writer12.String())

	// Création d'encodages JSON entièrement personnalisés
	dp8 := DiscountedProduct5{
		Product:  &kayak,
		Discount: 10.50,
	}
	dp8Json, err := dp8.MarshalJSON()
	if err == nil {
		Printfln("Write 13 : %v", dp8Json)
	}
}
