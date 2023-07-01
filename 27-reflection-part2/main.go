package main

import (
	"reflect"
	"strings"
)

/**
La fonction et la méthode du package reflect pour les pointeurs :
- PtrTo(type) : cette fonction renvoie un Type (reflect.Type) qui est un pointeur vers le Type reçu en argument.
- Elem() : cette méthode, qui est appelée sur un type pointeur (reflect.Type), renvoie le Type sous-jacent.
           Cette méthode panique lorsqu'elle est utilisée sur des types non pointeurs.

Les méthodes de reflect.Value pour travailler avec des types de pointeur :
- Addr() : cette méthode renvoie une valeur qui est un pointeur vers la valeur sur laquelle elle est appelée.
           Cette méthode panique si la méthode CanAddr renvoie false.
- CanAddr() : cette méthode renvoie vrai si la valeur peut être utilisée avec la méthode Addr.
- Elem() : cette méthode suit un pointeur et renvoie sa valeur. Cette méthode panique si elle est appelée sur une valeur non pointeur.
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
}
