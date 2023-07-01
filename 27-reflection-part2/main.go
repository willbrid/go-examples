package main

import (
	"reflect"
	"strings"
)

/**
La fonction et la méthode du package reflect pour les pointeurs :
- PtrTo(type) : cette fonction renvoie un Type (reflect.Type) qui est un pointeur vers le Type reçu en argument.
- Elem() : --- cette méthode, qui est appelée sur un type pointeur (reflect.Type), renvoie le Type sous-jacent.
           Cette méthode panique lorsqu'elle est utilisée sur des types non pointeurs.
		   --- Cette méthode renvoie le Type pour les éléments de tableau ou de tranche.
- Len() : cette méthode renvoie la longueur d'un type tableau. Cette méthode paniquera si elle est appelée sur d'autres types, y compris les tranches.
- ArrayOf(len, type) : cette fonction renvoie un Type qui décrit un tableau avec la taille et le type d'élément spécifiés.
- SliceOf(type) : cette fonction renvoie un Type qui décrit un tableau avec le type d'élément spécifié.

Les méthodes de reflect.Value pour travailler avec des types de pointeur, des types array ou slice :
- Addr() : cette méthode renvoie une valeur qui est un pointeur vers la valeur sur laquelle elle est appelée.
           Cette méthode panique si la méthode CanAddr renvoie false.
- CanAddr() : cette méthode renvoie vrai si la valeur peut être utilisée avec la méthode Addr.
- Elem() : cette méthode suit un pointeur et renvoie sa valeur. Cette méthode panique si elle est appelée sur une valeur non pointeur.
- Index(index) : cette méthode renvoie une valeur qui représente l'élément à l'index spécifié.
- Len() : cette méthode renvoie la longueur du tableau ou de la tranche.
- Cap() : cette méthode renvoie la capacité du tableau ou de la tranche.
- SetLen() : cette méthode définit la longueur d'une tranche. Elle ne peut pas être utilisé sur des tableaux.
- SetCap() : cette méthode définit la capacité d'une tranche. Elle ne peut pas être utilisé sur des tableaux.
- Slice(lo, hi) : cette méthode crée une nouvelle tranche avec les valeurs low et high spécifiées.
- Slice3(lo, hi, max) : cette méthode crée une nouvelle tranche avec les valeurs low, high et max spécifiées.
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
}
