package main

import (
	"reflect"
	"strings"
)

/**
Les fonctions et les méthodes du package reflect pour les pointeurs ou les maps :
- PtrTo(type) : cette fonction renvoie un Type (reflect.Type) qui est un pointeur vers le Type reçu en argument.
- Elem() : --- cette méthode, qui est appelée sur un type pointeur (reflect.Type), renvoie le Type sous-jacent.
           Cette méthode panique lorsqu'elle est utilisée sur des types non pointeurs.
		   --- Cette méthode renvoie reflect.Type pour les éléments de tableau ou de tranche.
		   --- Cette méthode renvoie reflect.Type pour les valeurs de map.
- Len() : cette méthode renvoie la longueur d'un type tableau. Cette méthode paniquera si elle est appelée sur d'autres types, y compris les tranches.
- ArrayOf(len, type) : cette fonction renvoie un Type qui décrit un tableau avec la taille et le type d'élément spécifiés.
- SliceOf(type) : cette fonction renvoie un Type qui décrit un tableau avec le type d'élément spécifié.
- Key() : cette méthode renvoie le Type (reflect.Type) des clés de la map.
- MapOf(keyType, valueType) : pour les maps, cette fonction renvoie un nouveau reflect.Type qui reflète le type de mappage avec
							  les types de clé et de valeur spécifiés, tous deux décrits à l'aide d'un reflect.Type.
- MakeMap(type) : pour les maps, cette fonction renvoie un reflect.Value qui reflète une map créée avec le type spécifié.
- MakeMapWithSize(type, size) : pour les maps, cette fonction renvoie un reflect.Value qui reflète une map créée avec le type et la taille spécifiés.
- StructOf(fields) : pour les classes (struct), cette fonction crée un nouveau type de struct, en utilisant la tranche reflect.StructField spécifiée
					 pour définir les champs. Seuls les champs exportés peuvent être spécifiés.

Les méthodes de reflect.Type pour travailler avec les classes (struct) :
- NumField() : cette méthode renvoie le nombre de champs définis par le type struct.
- Field(index) : cette méthode renvoie le champ à l'index spécifié, représenté par un reflect.StructField.
- FieldByIndex(indices) : cette méthode accepte une tranche int, qu'elle utilise pour localiser un champ imbriqué, qui est représenté par un reflect.StructField.
- FieldByName(name) : cette méthode renvoie le champ avec le nom spécifié, qui est représenté par un reflect.StructField.
                      Les résultats sont un reflect.StructField qui représente le champ et un bool qui indique si une correspondance a été trouvée.
- FieldByNameFunc(func) : cette méthode transmet le nom de chaque champ inclus des champs imbriqués à la fonction spécifiée et
                          renvoie le premier champ pour lequel la fonction renvoie true. Les résultats sont un StructField qui représente le champ
						  et un bool qui indique si une correspondance a été trouvée.

Les attributs de la classe reflect.StructField :
--- Name : ce champ stocke le nom du champ reflété.
--- PkgPath : ce champ renvoie le nom du package, qui est utilisé pour déterminer si un champ a été exporté.
              Pour les champs reflétés qui sont exportés, ce champ renvoie la chaîne vide. Pour les champs reflétés qui n'ont pas été exportés,
			  ce champ renvoie le nom du package, qui est le seul package dans lequel le champ peut être utilisé.
--- Type : ce champ renvoie le type reflété du champ reflété, décrit à l'aide d'un Type.
--- Tag : ce champ renvoie le tag struct associée au champ reflété.
--- Index : ce champ renvoie une tranche int, qui indique l'index du champ utilisé par la méthode FieldByIndex.
--- Anonyme : ce champ renvoie true si le champ reflété est intégré et false sinon.

Les méthodes définies par le type reflect.StructTag :
-- Get(key) : cette méthode renvoie une chaîne contenant la valeur de la clé spécifiée ou la chaîne vide si aucune valeur n'a été définie.
-- Lookup(key) : cette méthode renvoie une chaîne contenant la valeur de la clé spécifiée ou la chaîne vide si aucune valeur n'a été définie,
                et un booléen qui est vrai si la valeur a été définie et faux sinon.


Les méthodes de reflect.Value pour travailler avec des types de pointeur, des types array ou slice, des types map :
- Addr() : cette méthode renvoie une valeur qui est un pointeur vers la valeur sur laquelle elle est appelée.
           Cette méthode panique si la méthode CanAddr renvoie false.
- CanAddr() : cette méthode renvoie vrai si la valeur peut être utilisée avec la méthode Addr.
- Elem() : cette méthode suit un pointeur et renvoie sa valeur. Cette méthode panique si elle est appelée sur une valeur non pointeur.
- Index(index) : cette méthode renvoie une valeur qui représente l'élément à l'index spécifié.
- Len() : cette méthode renvoie la longueur du tableau ou de la tranche.
          pour les maps, cette méthode renvoie le nombre de paires clé-valeur contenues dans une map.
- Cap() : cette méthode renvoie la capacité du tableau ou de la tranche.
- SetLen() : cette méthode définit la longueur d'une tranche. Elle ne peut pas être utilisé sur des tableaux.
- SetCap() : cette méthode définit la capacité d'une tranche. Elle ne peut pas être utilisé sur des tableaux.
- Slice(lo, hi) : cette méthode crée une nouvelle tranche avec les valeurs low et high spécifiées.
- Slice3(lo, hi, max) : cette méthode crée une nouvelle tranche avec les valeurs low, high et max spécifiées.
- MapKeys() : cette méthode renvoie un []reflect.Value, contenant les clés de la map.
- MapIndex(key) : cette méthode renvoie la valeur qui correspond à la clé spécifiée, qui est également exprimée sous forme de reflect.Value.
                  La valeur zéro est renvoyée si la map ne contient pas la clé spécifiée, qui peut être détectée
				  en appelant la méthode IsValid, qui renverra false.
- MapRange() : cette méthode renvoie un *MapIter, qui permet d'itérer le contenu de la map.
- SetMapIndex(key, val) : cette méthode définit la clé et la valeur spécifiées, toutes deux exprimées à l'aide de l'interface reflect.Value.


Fonctions d'ajout d'éléments aux tranches du package reflect -> ces fonctions acceptent les arguments reflect.Type ou reflect.Value :
- MakeSlice(type, len, cap) : cette fonction crée un reflect.Value qui reflète une nouvelle tranche, en utilisant un reflect.Type pour désigner
                              le type d'élément et avec la longueur et la capacité spécifiées.
- Append(sliceVal, ...val) : cette fonction ajoute une ou plusieurs valeurs à la tranche spécifiée, qui sont toutes exprimées à l'aide
                             de l'interface Value. Le résultat est la tranche modifiée. La fonction panique lorsqu'elle est utilisée sur
							 tout type de reflect.Type autre qu'une tranche ou si les types des valeurs ne correspondent pas au type d'élément de tranche.
- AppendSlice(sliceVal, sliceVal) : cette fonction ajoute une tranche à une autre. La fonction panique si l'une des valeurs ne représente pas
                                    une tranche ou si les types de tranche ne sont pas compatibles.
- Copy(dst, src) : cette fonction copie les éléments de la tranche ou du tableau reflété par la valeur src vers la tranche ou le tableau reflété
                   par la valeur dst. Les éléments sont copiés jusqu'à ce que la tranche de destination soit pleine ou que tous
				   les éléments source aient été copiés. La source et la destination doivent avoir le même type d'élément.


Les méthodes définies par la classe reflect.MapIter :
- Next() : cette méthode passe à la paire clé-valeur suivante dans la map. Le résultat de cette méthode est un booléen indiquant
          s'il existe d'autres paires clé-valeur à lire. Cette méthode doit être appelée avant la méthode Key ou Value.
- Key() : cette méthode renvoie la valeur représentant la clé de map à la position actuelle.
- Value() : cette méthode renvoie la valeur représentant la valeur de la map à la position actuelle.
**/

func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PtrTo(t)
}

func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}

	return t
}

var stringPtrType = reflect.TypeOf((*string)(nil))

/*
*
La fonction transformString identifie les valeurs *string et utilise la méthode reflect.Value.Elem pour obtenir la valeur de chaîne
afin qu'elle puisse être transmise à la fonction strings.ToUpper
*
*/
func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}

/*
*
Le checkElemType utilise la méthode Kind pour identifier les tableaux et les tranches et utilise la méthode Elem pour obtenir le Type des éléments.
Ceux-ci sont comparés au type du premier paramètre pour voir si la valeur peut être ajoutée en tant qu'élément.
*
*/
func checkElemType(val interface{}, arrOrSlice interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)
	return (arrOrSliceType.Kind() == reflect.Array || arrOrSliceType.Kind() == reflect.Slice) && arrOrSliceType.Elem() == elemType
}

/*
*
Lors des modifications sur un tableau (array), ce tableau a besoin d'être référencé via un pointeur.
*
*/
func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr && arrayOrSliceVal.Elem().Kind() == reflect.Array && arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

/*
*
La fonction enumerateStrings vérifie le résultat Kind pour s'assurer qu'il s'agit d'un tableau ou d'une tranche de chaînes.
Il est facile de ne pas savoir quelle méthode Elem est utilisée dans ce processus car Type et Value définissent les méthodes Kind et Elem.
Les méthodes Kind effectuent la même tâche, mais appeler la méthode Elem sur une tranche ou un tableau reflect.Value provoque une panique, tandis que
l'appel de la méthode Elem sur une tranche ou un tableau reflect.Type renvoie le Type des éléments.
*
*/
func enumerateStrings(arrayOrSlice interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	if (arrayOrSliceVal.Kind() == reflect.Array || arrayOrSliceVal.Kind() == reflect.Slice) && arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
		for i := 0; i < arrayOrSliceVal.Len(); i++ {
			Printfln("Element : %v, Value : %v", i, arrayOrSliceVal.Index(i).String())
		}
	}
}

/*
*
La fonction findAndSplit énumère la tranche, en recherchant l'élément spécifié, ce qui est fait à l'aide de la méthode Interface,
qui permet de comparer les éléments de la tranche sans avoir à traiter de types spécifiques. Une fois l'élément cible localisé,
la méthode Slice est utilisée pour créer et renvoyer une nouvelle tranche.
*
*/
func findAndSplit(slice interface{}, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				return sliceVal.Slice(0, i+1)
			}
		}
	}

	return slice
}

/*
*
La fonction pickValues crée une nouvelle tranche à l'aide de reflect.Type à partir d'une tranche existante et utilise
la fonction Append pour ajouter des valeurs à la nouvelle tranche.
*
*/
func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice, sliceVal.Index(index))
		}
		return newSlice
	}

	return nil
}

func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)
	if mapType.Kind() == reflect.Map {
		Printfln("Key type: %v, Val type: %v", mapType.Key(), mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key : %v, Value : %v", keyVal, reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}

/*
*
Il est important d'appeler la méthode Next avant d'appeler les méthodes Key et Value et d'éviter d'appeler ces méthodes
lorsque la méthode Next renvoie false.
*
*/
func printMapContentsWithMapIter(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		iter := mapValue.MapRange()
		for iter.Next() {
			Printfln("Map Key : %v, Value : %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

func setMap(m interface{}, key interface{}, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)
	if mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type() && mapValue.Type().Elem() == valValue.Type() {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map")
	}
}

func removeFromMap(m interface{}, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	if mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type() {
		/**
		La méthode SetMapIndex supprimera une clé de la map si l'argument de valeur est la valeur zéro pour le type de valeur de map.
		C'est un problème lorsqu'il s'agit de types intégrés, tels que int et float64, où la valeur zéro est une entrée de map valide.
		Pour empêcher SetMapIndex de définir des valeurs sur zéro, la fonction removeFromMap crée une instance de la classe reflect.Value.
		**/
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}

/*
*
La fonction createMap accepte une tranche de valeurs et une fonction. La tranche est énumérée et la fonction est appelée sur chaque élément,
avec les valeurs d'origine et transformées utilisées pour remplir une map, qui est renvoyée comme résultat de la fonction.
*
*/
func createMap(slice interface{}, op func(interface{}) interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		mapType := reflect.MapOf(sliceVal.Type().Elem(), sliceVal.Type().Elem())
		mapVal := reflect.MakeMap(mapType)
		for i := 0; i < sliceVal.Len(); i++ {
			elemVal := sliceVal.Index(i)
			mapVal.SetMapIndex(elemVal, reflect.ValueOf(op(elemVal.Interface())))
		}
		return mapVal.Interface()
	}

	return nil
}

/*
*
La fonction inspectStructs définit un paramètre variadique à travers lequel elle reçoit des valeurs.
La fonction TypeOf est utilisée pour obtenir le type reflété et la méthode Kind est utilisée pour confirmer que chaque type est une structure.
Le Type reflété est transmis à la fonction inspectStructType, dans laquelle la méthode NumField est utilisée dans une boucle for,
ce qui permet d'énumérer les champs structs à l'aide de la méthode Field.

La même approche peut être utilisée pour inspecter les champs qui sont des pointeurs vers des types struct, avec l'utilisation de la méthode Type.Elem
pour obtenir le type auquel le pointeur fait référence.
*
*/
func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType(structType)
		}
	}
}

func inspectStructType(structType reflect.Type) {
	Printfln("--- Struct Type : %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		Printfln("Field %v: Name : %v, Type : %v, Exported : %v", field.Index, field.Name, field.Type, field.PkgPath == "")
	}
	Printfln("--- End Struct Type : %v", structType)
}

func inspectStructsImprov(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructTypeImprov([]int{}, structType)
		}
	}
}

func inspectStructTypeImprov(baseIndex []int, structType reflect.Type) {
	Printfln("--- Struct Type : %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := structType.Field(i)
		Printfln("Field %v: Name : %v, Type : %v, Exported : %v", field.Index, field.Name, field.Type, field.PkgPath == "")
		if field.Type.Kind() == reflect.Struct {
			field := structType.FieldByIndex(fieldIndex)
			inspectStructTypeImprov(fieldIndex, field.Type)
		}
	}
	Printfln("--- End Struct Type : %v", structType)
}

/*
*
La fonction describeField utilise la méthode FieldByName, qui localise le premier champ avec le nom spécifié et renvoie un StructField avec
un champ Index correctement défini. Une boucle for est utilisée pour remonter la hiérarchie des types, en examinant tour à tour chaque parent.
*
*/
func describeField(s interface{}, fieldName string) {
	structType := reflect.TypeOf(s)
	field, found := structType.FieldByName(fieldName)

	if found {
		Printfln("Found : %v, Type : %v, Index : %v", field.Name, field.Type, field.Index)
		index := field.Index
		for len(index) > 1 {
			index = index[0 : len(index)-1]
			field = structType.FieldByIndex(index)
			Printfln("Parent : %v, Type : %v, Index : %v", field.Name, field.Type, field.Index)
		}
		Printfln("Top-Level Type : %v", structType)
	} else {
		Printfln("Field %v not found", fieldName)
	}
}

/*
*
La fonction inspectTags énumère les champs définis par un type struct et utilise à la fois les méthodes Get et Lookup pour obtenir un tag spécifié.
La fonction est appliquée au type Person, qui définit la tag alias sur certains de ses champs.
*
*/
type Person struct {
	Name    string `alias:"id"`
	City    string `alias:""`
	Country string
}

func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v : %v", field.Name, tagName, valGet)
		Printfln("Field: %v, Tag %v : %v, Set : %v", field.Name, tagName, valLookup, ok)
	}
}

func main() {
	name1 := "Alice"
	t := reflect.TypeOf(name1)
	Printfln("Original : Type : %v", t)
	pt := createPointerType(t)
	Printfln("Pointer Type : %v", pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))

	name2 := "Alice"
	transformString(&name2)
	Printfln("Follow pointer value: %v", name2)

	name3 := "Alice"
	city3 := "London"
	hobby3 := "Running"
	slice3 := []string{name3, city3, hobby3}
	array3 := [3]string{name3, city3, hobby3}
	Printfln("Slice (string) : %v", checkElemType("testString", slice3))
	Printfln("Array (string) : %v", checkElemType("testString", array3))
	Printfln("Array (int) : %v", checkElemType(10, array3))

	name4 := "Alice"
	city4 := "London"
	hobby4 := "Running"
	slice4 := []string{name4, city4, hobby4}
	array4 := [3]string{name4, city4, hobby4}
	Printfln("Original slice : %v", slice4)
	newCity := "Paris"
	setValue(slice4, 1, newCity)
	Printfln("Modified slice : %v", slice4)
	Printfln("Original array: %v", array4)
	newCity = "Rome"
	setValue(&array4, 1, newCity)
	Printfln("Modified array : %v", array4)
	enumerateStrings(slice4)
	enumerateStrings(array4)

	name5 := "Alice"
	city5 := "London"
	hobby5 := "Running"
	slice5 := []string{name5, city5, hobby5}
	Printfln("Strings: %v", findAndSplit(slice5, "London"))
	numbers5 := []int{1, 3, 4, 5, 7}
	Printfln("Numbers: %v", findAndSplit(numbers5, 4))

	name6 := "Alice"
	city6 := "London"
	hobby6 := "Running"
	slice6 := []string{name6, city6, hobby6, "Bob", "Paris", "Soccer"}
	picked := pickValues(slice6, 0, 3, 5)
	Printfln("Picked values: %v", picked)

	pricesMap := map[string]float64{
		"Kayak":       279,
		"Lifejacket":  48.95,
		"Soccer Ball": 19.50,
	}
	describeMap(pricesMap)
	printMapContents(pricesMap)
	printMapContentsWithMapIter(pricesMap)
	setMap(pricesMap, "Kayak", 100.00)
	setMap(pricesMap, "Hat", 10.00)
	removeFromMap(pricesMap, "Lifejacket")
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	reverse := func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return strings.ToUpper(str)
		}
		return val
	}
	namesMap := createMap(names, reverse).(map[string]string)
	for k, v := range namesMap {
		Printfln("Key : %v, Value : %v", k, v)
	}

	/**
	Inspection d'une classe struct
	**/
	inspectStructs(Purchase{})
	inspectStructsImprov(Purchase{}) // Inspection des champs imbriqués d'une classe struct
	describeField(Purchase{}, "Price")
	inspectTags(Person{}, "alias") // Inspection des tags d'une classe

	/**
	Cet exemple crée une classe qui a les mêmes caractéristiques que la classe Person, avec les champs Name, City et Country.
	Les champs sont décrits en créant des valeurs reflect.StructField, qui ne sont que des classes Go normales.
	La fonction New est utilisée pour créer une nouvelle valeur (reflect.Value) à partir de la classe, qui est transmise à la fonction inspectTags
	**/
	stringType := reflect.TypeOf("this is a string")
	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})
	inspectTags(reflect.New(structType), "alias")
}
