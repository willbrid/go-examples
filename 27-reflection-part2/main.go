package main

import (
	"reflect"
	"strings"
)

/**
La fonction et la méthode du package reflect pour les pointeurs ou les maps :
- PtrTo(type) : cette fonction renvoie un Type (reflect.Type) qui est un pointeur vers le Type reçu en argument.
- Elem() : --- cette méthode, qui est appelée sur un type pointeur (reflect.Type), renvoie le Type sous-jacent.
           Cette méthode panique lorsqu'elle est utilisée sur des types non pointeurs.
		   --- Cette méthode renvoie reflect.Type pour les éléments de tableau ou de tranche.
		   --- Cette méthode renvoie reflect.Type pour les valeurs de map.
- Len() : cette méthode renvoie la longueur d'un type tableau. Cette méthode paniquera si elle est appelée sur d'autres types, y compris les tranches.
- ArrayOf(len, type) : cette fonction renvoie un Type qui décrit un tableau avec la taille et le type d'élément spécifiés.
- SliceOf(type) : cette fonction renvoie un Type qui décrit un tableau avec le type d'élément spécifié.
- Key() : cette méthode renvoie le Type (reflect.Type) des clés de la map.

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

La fonction reflect pour créer des types de map :
- MapOf(keyType, valueType) : cette fonction renvoie un nouveau reflect.Type qui reflète le type de mappage avec les types de clé et de valeur spécifiés,
                               tous deux décrits à l'aide d'un reflect.Type.

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
}
