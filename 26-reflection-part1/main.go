package main

import (
	"fmt"
	"reflect"
	"strings"
)

/**
L'interface vide permet à la fonction printDetailsWithInterface de recevoir n'importe quel type mais n'autorise pas l'accès à des fonctionnalités
spécifiques car l'interface ne définit aucune méthode. Une assertion de type est nécessaire pour restreindre l'interface vide à un
type spécifique, ce qui permet ensuite de traiter chaque valeur.

La limitation de cette approche est que la fonction printDetailsWithInterface ne peut traiter que des types connus à l'avance.
Chaque fois que nous ajoutons un type au projet, nous devons étendre la fonction printDetailsWithInterface pour gérer ce type.
De nombreux projets traiteront d'un ensemble suffisamment petit de types pour que cela ne soit pas un problème ou pourront définir des interfaces
avec des méthodes qui donnent accès à des fonctionnalités communes. Reflection résout ce problème pour les projets pour lesquels
ce n'est pas le cas, soit parce qu'il y a un grand nombre de types à traiter, soit parce que les interfaces et les méthodes ne peuvent pas être écrites.

Le point clé à retenir est qu'il y a deux aspects de la réflexion qui fonctionnent ensemble : le type reflété et la valeur reflétée.
Le type reflété nous donne accès aux détails d'un type Go sans savoir à l'avance de quoi il s'agit. Nous pouvons explorer le type reflété, en explorant
ses détails et ses caractéristiques à travers les méthodes définies par l'interface Type.
La valeur reflétée nous permet de travailler avec la valeur spécifique qui nous a été fournie. Nous ne pouvons pas simplement lire un champ struct
ou appeler une méthode, par exemple, comme nous le ferions dans un code normal lorsque nous ne savons pas à quel type nous avons affaire.
L'utilisation du type et de la valeur reflétés conduit à la verbosité du code. Si nous savons que nous avons affaire à une classe Product,
par exemple, nous pouvons simplement lire le champ Name et obtenir un résultat de chaîne. Si nous ne savons pas quel type est utilisé,
nous devons utiliser le type reflété pour déterminer si nous avons affaire à une classe et si elle a un champ Name.
Une fois que nous avons déterminé qu'il existe un tel champ, nous utilisons la valeur reflétée pour lire ce champ et obtenir sa valeur.

L'interface reflect.Type fournit des détails de base sur un type à travers les méthodes :
- Name() : cette méthode renvoie le nom du type.
- PkgPath() : cette méthode renvoie le chemin du package pour le type. La chaîne vide est renvoyée pour les types intégrés, tels que int et bool.
- Kind() : cette méthode renvoie le genre du type, en utilisant une valeur qui correspond à l'une des valeurs constantes définies par le package reflect.
- String() : cette méthode renvoie une représentation sous forme de chaîne du nom du type, y compris le nom du package.
- Comparable() : cette méthode renvoie true si les valeurs de ce type peuvent être comparées à l'aide de l'opérateur de comparaison standard.
- AssignableTo(type) : cette méthode renvoie true si des valeurs de ce type peuvent être affectées à des variables ou à des champs du type reflété spécifié.

Les constantes liées à la méthode reflect.Type.kind() :
- Bool : cette valeur correspond à un booléen.
- Int, Int8, Int16, Int32, Int64 : ces valeurs indiquent les différentes tailles des types entiers.
- Uint, Uint8, Uint16, Uint32, Uint64 : ces valeurs indiquent les différentes tailles des types entiers non signés.
- Float32, Float64 : ces valeurs indiquent les différentes tailles des types à virgule flottante.
- String : cette valeur désigne une chaîne.
- Struct : cette valeur désigne une classe.
- Array : cette valeur désigne un tableau.
- Slice : cette valeur indique une tranche.
- Map : cette valeur désigne un tableau associatif.
- Chan : cette valeur désigne un canal.
- Func : cette valeur définit une fonction.
- Interface : cette valeur désigne une interface.
- Ptr : cette valeur désigne un pointeur
- Uintptr : cette valeur indique un pointeur non sécurisé
**/

type Payment struct {
	Currency string
	Amount   float64
}

func printDetails(values ...Product) {
	for _, elem := range values {
		Printfln("Product: Name: %v, Category: %v, Price: %v", elem.Name, elem.Category, elem.Price)
	}
}

func printDetailsWithInterface(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v", val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func printDetailsWithReflection(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		// La fonction TypeOf renvoie le type reflété (reflect.Type), qui est décrit par l'interface Type.
		elemType := reflect.TypeOf(elem)
		// La fonction ValueOf renvoie la valeur reflétée (reflect.Value), qui est représentée par l'interface Value.
		elemValue := reflect.ValueOf(elem)
		// la méthode Type.Kind() renvoie le type est en cours de traitement.
		// Le package reflect définit des constantes qui identifient les différents types de type dans Go.
		if elemType.Kind() == reflect.Struct {
			// La méthode Type.NumField() renvoie le nombre de champs d'une classe
			for i := 0; i < elemType.NumField(); i++ {
				/**
				L'appel de la méthode Type.Field() sur un type reflété renvoie une classe reflect.StructField, qui décrit un seul champ, y compris un champ Name.
				L'appel de la méthode Type.Field() sur la valeur reflétée renvoie une structure reflect.Value, qui représente la valeur du champ.
				**/
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v : %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetailsWithReflectionAndTypePath(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		Printfln("Name: %v, PkgPath: %v, Kind: %v", elemType.Name(), getTypePath(elemType), elemType.Kind())
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v : %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

func main() {
	product1 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	printDetails(product1)

	product2 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer2 := Customer{Name: "Alice", City: "New Yorl"}
	printDetailsWithInterface(product2, customer2)

	product3 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer3 := Customer{Name: "Alice", City: "New Yorl"}
	payment3 := Payment{Currency: "USD", Amount: 100.50}
	printDetailsWithReflection(product3, customer3, payment3, 10, true)

	product4 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer4 := Customer{Name: "Alice", City: "New Yorl"}
	payment4 := Payment{Currency: "USD", Amount: 100.50}
	printDetailsWithReflectionAndTypePath(product4, customer4, payment4, 10, true)
}
