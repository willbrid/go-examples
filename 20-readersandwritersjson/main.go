package main

import (
	"encoding/json"
	"io"
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

	// Création d'encodage JSON entièrement personnalisé
	dp8 := DiscountedProduct5{
		Product:  &kayak,
		Discount: 10.50,
	}
	dp8Json, err := dp8.MarshalJSON()
	if err == nil {
		Printfln("Write 13 : %v", dp8Json)
		Printfln("Write 13 : %v", string(dp8Json)) // Conversion du byte en string
	}

	// La fonction constructeur NewDecoder crée un décodeur, qui peut être utilisé pour décoder les données JSON obtenues à partir d'un reader.
	/**
	Un lecteur est créé, qui produira des données à partir d'une chaîne contenant une séquence de valeurs, séparées par des espaces
	(la spécification JSON permet de séparer les valeurs par des espaces ou des caractères de saut de ligne).
	La première étape du décodage des données consiste à créer le décodeur, qui accepte un reader. L'on veut décoder plusieurs valeurs,
	nous appelons donc la méthode Decode dans une boucle for. Le décodeur est capable de sélectionner le type de données Go approprié pour les valeurs JSON,
	et cela est réalisé en fournissant un pointeur vers une interface vide comme argument de la méthode Decode.
	La méthode Decode renvoie une erreur, qui indique des problèmes de décodage mais est également utilisée pour signaler la fin des données
	à l'aide de l'erreur io.EOF. Une boucle for décode à plusieurs reprises les valeurs jusqu'à EOF, puis nous utilisons une autre boucle for
	pour écrire chaque type et valeur décodés à l'aide des verbes de formatage.
	**/
	reader1 := strings.NewReader(`true "Hello" 99.99 200`)
	val1s := []interface{}{}
	decoder1 := json.NewDecoder(reader1)
	for {
		var decodedVal interface{}
		/**
		Cette méthode decoder1.Decode lit et décode les données, qui sont utilisées pour créer la valeur spécifiée. La méthode renvoie une erreur qui indique
		des problèmes de décodage des données vers le type requis ou EOF.
		Decode lit la prochaine valeur encodée en JSON à partir de son entrée et la stocke dans la valeur pointée par decodedVal.
		**/
		err := decoder1.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error : %v", err.Error())
			}
			break
		}
		val1s = append(val1s, decodedVal)
	}
	/**
	Cette syntaxe val.(json.Number) est utilisée en Go pour convertir une valeur de type interface{} en un type plus spécifique.
	Dans ce cas, la valeur est convertie en json.Number.
	En Go, l'interface{} est un type générique qui peut contenir des valeurs de n'importe quel type.
	**/
	for _, val := range val1s {
		if num, ok := val.(json.Number); ok {
			// Cette méthode num.Int64 renvoie la valeur décodée sous la forme d'un int64 et une erreur qui indique si la valeur ne peut pas être convertie.
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
				// Cette méthode num.Float64 renvoie la valeur décodée sous la forme d'un float64 et une erreur qui indique si la valeur ne peut pas être convertie.
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				// Cette méthode num.String renvoie la chaîne non convertie à partir des données JSON.
				Printfln("Decoded String: %v", num.String())
			}
		}
		Printfln("Decoded (%T): %v", val, val)
	}

	/**
	Le décodeur renverra une erreur s'il ne peut pas décoder une valeur JSON dans un type spécifié.
	Cette technique doit être utilisée uniquement lorsque nous sommes sûr de comprendre les données JSON qui seront décodées.
	**/
	reader2 := strings.NewReader(`true "Hello" 99.99 200`)
	var (
		bval2  bool
		sval2  string
		fpval2 float64
		ival2  int
	)
	val2s := []interface{}{&bval2, &sval2, &fpval2, &ival2}
	decoder2 := json.NewDecoder(reader2)
	for i := 0; i < len(val2s); i++ {
		// Decode lit la prochaine valeur encodée en JSON à partir de son entrée et la stocke dans la valeur pointée par val2s[i] (qui est un pointeur).
		err := decoder2.Decode(val2s[i])
		if err != nil {
			Printfln("Error : %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval2, bval2)
	Printfln("Decoded (%T): %v", sval2, sval2)
	Printfln("Decoded (%T): %v", fpval2, fpval2)
	Printfln("Decoded (%T): %v", ival2, ival2)

	/**
	Les données JSON source contiennent deux tableaux, dont l'un ne contient que des nombres et l'autre mélange des nombres et des chaînes.
	Le décodeur n'essaie pas de déterminer si un tableau JSON peut être représenté à l'aide d'un seul type Go et décode chaque tableau en
	une tranche d'interface vide.
	**/
	reader3 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	val3s := []interface{}{}
	decoder3 := json.NewDecoder(reader3)
	for {
		var decodedVal interface{}
		err := decoder3.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		val3s = append(val3s, decodedVal)
	}
	for _, val := range val3s {
		Printfln("Decoded (%T): %v", val, val)
	}

	/**
	Dans l'exemple ci-dessus, chaque valeur est typée en fonction de la valeur JSON, mais le type de la tranche est l'interface vide.
	Si nous connaissons à l'avance la structure des données JSON et que nous décodons un tableau contenant un seul type de données JSON,
	nous pouvons alors passer une tranche Go du type souhaité à la méthode Decode.
	**/
	reader4 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	ints := []int{}
	mixed := []interface{}{}
	val4s := []interface{}{&ints, &mixed}
	decoder4 := json.NewDecoder(reader4)
	for i := 0; i < len(val4s); i++ {
		// Decode lit la prochaine valeur encodée en JSON à partir de son entrée et la stocke dans la valeur pointée par val2s[i] (qui est un pointeur).
		err := decoder4.Decode(val4s[i])
		if err != nil {
			Printfln("Error : %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)

	/**
	Les objets JavaScript sont exprimés sous forme de paires clé-valeur, ce qui facilite leur décodage en cartes Go.
	L'approche la plus sûre consiste à définir une map avec des clés de chaîne et des valeurs d'interface vides,
	ce qui garantit que toutes les paires clé-valeur dans les données JSON peuvent être décodées dans la map.
	**/
	reader5 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m5 := map[string]interface{}{}
	decoder5 := json.NewDecoder(reader5)
	err5 := decoder5.Decode(&m5)
	if err5 != nil {
		Printfln("Error: %v", err5.Error())
	} else {
		Printfln("Map: %T, %v", m5, m5)
		for k, v := range m5 {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	/**
	Un seul objet JSON peut être utilisé pour plusieurs types de données en tant que valeurs, mais si nous savons à l'avance que
	nous allons décoder un objet JSON qui a un seul type de valeur, nous pouvons être plus précis lors de la définition de la map
	dans laquelle les données seront être décodé
	**/
	reader6 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m6 := map[string]float64{}
	decoder6 := json.NewDecoder(reader6)
	err6 := decoder6.Decode(&m6)
	if err6 != nil {
		Printfln("Error: %v", err6.Error())
	} else {
		Printfln("Map: %T, %v", m6, m6)
		for k, v := range m6 {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	/**
	Le décodeur décode l'objet JSON et utilise les clés pour définir les valeurs des champs de structure exportés. La capitalisation des champs et
	des clés JSON ne doit pas nécessairement correspondre, et le décodeur ignorera toute clé JSON pour laquelle il n'y a pas de champ struct et ignorera
	tout champ struct pour lequel il n'y a pas de clé JSON. Les objets JSON contiennent des majuscules différentes et ont plus ou moins de clés
	que la structure Product des champs.
	**/
	reader7 := strings.NewReader(`
		{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports"}
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true}
	`)
	decoder7 := json.NewDecoder(reader7)
	for {
		var val Product
		err := decoder7.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error : %v", err.Error())
			}
			break
		} else {
			Printfln("Name : %v, Category : %v, Price : %v", val.Name, val.Category, val.Price)
		}
	}

	/**
	Par défaut, le décodeur ignorera les clés JSON pour lesquelles il n'y a pas de champ de structure correspondant.
	Ce comportement peut être modifié en appelant la méthode DisallowUnknownFields, qui déclenche une erreur lorsqu'une telle clé est rencontrée.
	**/
	reader8 := strings.NewReader(`
		{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports"}
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true}
	`)
	decoder8 := json.NewDecoder(reader8)
	// DisallowUnknownFields fait que Decoder renvoie une erreur lorsque la destination est une structure et que l'entrée contient des
	// clés d'objet qui ne correspondent à aucun champ exporté non ignoré dans la destination.
	decoder8.DisallowUnknownFields()
	for {
		var val Product
		err := decoder8.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error : %v", err.Error())
			}
			break
		} else {
			Printfln("Name : %v, Category : %v, Price : %v", val.Name, val.Category, val.Price)
		}
	}

	/**
	La balise appliquée au champ Discount indique au décodeur que la valeur de ce champ doit être obtenue à partir de la
	clé JSON nommée offer et que la valeur sera analysée à partir d'une chaîne, au lieu du numéro JSON qui serait généralement
	attendu pour un Go float64 valeur.
	**/
	reader9 := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	decoder9 := json.NewDecoder(reader9)
	for {
		var val DiscountedProduct6
		err := decoder9.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error : %v", err.Error())
			}
			break
		} else {
			Printfln("Name : %v, Category : %v, Price : %v, Discount: %v", val.Name, val.Category, val.Price, val.Discount)
		}
	}

	// Création de décodage JSON entièrement personnalisé
	data10 := []byte(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	var dp10 DiscountedProduct7
	dp10.UnmarshalJSON(data10)
	if err == nil {
		Printfln("Name : %v, Category : %v, Price : %v, Discount: %v", dp10.Name, dp10.Category, dp10.Price, dp10.Discount)
	}
}
