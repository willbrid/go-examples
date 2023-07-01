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

Les méthodes de la classe reflect.Value :
- Kind() : cette méthode renvoie le genre du type de la valeur.
- Type() : cette méthode renvoie le type de la valeur.
- IsNil() : cette méthode renvoie true si la valeur est nil. Cette méthode paniquera si la valeur sous-jacente n'est pas une fonction, une interface,
            un pointeur, une tranche ou un canal.
- IsZero() : cette méthode renvoie true si la valeur sous-jacente est la valeur zéro de son type.
- Bool() : cette méthode renvoie la valeur bool sous-jacente. La méthode panique si le Kind de la valeur sous-jacente n'est pas Bool.
- Bytes() : cette méthode renvoie la valeur []byte sous-jacente. La méthode panique si la valeur sous-jacente n'est pas une tranche d'octet.
- Int() : cette méthode renvoie la valeur sous-jacente sous la forme d'un int64. La méthode panique si le Kind de la valeur sous-jacente
          n'est pas Int, Int8, Int16, Int32 ou Int64.
- Uint() : cette méthode renvoie la valeur sous-jacente sous la forme d'un uint64. La méthode panique si le Kind de la valeur sous-jacente
           n'est pas Uint, Uint8, Uint16, Uint32 ou Uint64.
- Float() : cette méthode renvoie la valeur sous-jacente sous la forme d'un float64. La méthode panique si le Kind de la valeur sous-jacente
            n'est pas Float32 ou Float64.
- String() : cette méthode renvoie la valeur sous-jacente sous forme de chaîne si le Kind de la valeur est String. Pour les autres
             valeurs Kind, cette méthode renvoie la chaîne <T Value> où T est le type sous-jacent, tel que <int Value>.
- Elem() : cette méthode renvoie la valeur à laquelle un pointeur se réfère. Cette méthode peut également être utilisée avec des interfaces.
           Cette méthode panique si le Kind de la valeur sous-jacente n'est pas Ptr.
- IsValid() : cette méthode renvoie false si Value est la valeur zéro, créée en tant que Value{} plutôt qu'obtenue à l'aide de ValueOf, par exemple.
              Cette méthode ne concerne pas les valeurs reflétées qui sont la valeur zéro de leur type réfléchi.
			  Si cette méthode renvoie false, toutes les autres méthodes Value paniqueront.
- Interface() : cette méthode renvoie la valeur sous-jacente à l'aide de l'interface vide. Cette méthode paniquera si elle est utilisée
                sur des champs de structure non exportés.
- CanInterface() : cette méthode retourne true si la méthode Interface peut être utilisée sans paniquer.
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

func printDetailsWithReflectionAndReflectValueStruct(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		switch elemValue.Kind() {
		case reflect.Bool:
			var val bool = elemValue.Bool()
			Printfln("Bool: %v", val)
		case reflect.Int:
			var val int64 = elemValue.Int()
			Printfln("Int: %v", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = elemValue.Float()
			Printfln("Float: %v", val)
		case reflect.String:
			var val string = elemValue.String()
			Printfln("String: %v", val)
		case reflect.Ptr:
			var val reflect.Value = elemValue.Elem()
			if val.Kind() == reflect.Int {
				Printfln("Pointer to Int : %v", val.Int())
			}
		default:
			Printfln("Other: %v", elemValue.String())
		}
	}
}

/*
*
Cette technique commence par une valeur nulle et la convertit en un pointeur vers une valeur int, qui est ensuite transmise à la fonction TypeOf
pour obtenir un Type pouvant être utilisé dans les comparaisons. cette approche évite d'avoir à définir une variable juste pour obtenir son Type.
*
*/
var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func printDetailsWithReflectionByComparingType1(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)
		if elemType == intPtrType {
			Printfln("Pointer to Int: %v", elemValue.Elem().Int())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

func printDetailsWithReflectionByComparingType2(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)
		if elemType == intPtrType {
			Printfln("Pointer to Int: %v", elemValue.Elem().Int())
		} else if elemType == byteSliceType {
			Printfln("Byte slice: %v", elemValue.Bytes())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

/*
*
La fonction selectValue sélectionne une valeur dans une tranche sans connaître le type d'élément de la tranche. La valeur est extraite de la tranche à
l'aide de la méthode Index. Ce qui est important, c'est que la méthode Index renvoie une valeur, qui n'est utile que pour le code qui utilise la réflexion.
La méthode Interface est utilisée pour obtenir une valeur qui peut être utilisée comme résultat de la fonction.
*
*/
func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
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

	product5 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer5 := Customer{Name: "Alice", City: "New Yorl"}
	payment5 := Payment{Currency: "USD", Amount: 100.50}
	number5 := 100
	printDetailsWithReflectionAndReflectValueStruct(product5, customer5, payment5, true, 10, &number5, "Alice", 23.30)

	product6 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer6 := Customer{Name: "Alice", City: "New Yorl"}
	payment6 := Payment{Currency: "USD", Amount: 100.50}
	number6 := 100
	printDetailsWithReflectionByComparingType1(product6, customer6, payment6, true, 10, &number6, "Alice", 23.30)

	product7 := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer7 := Customer{Name: "Alice", City: "New Yorl"}
	payment7 := Payment{Currency: "USD", Amount: 100.50}
	number7 := 100
	slice7 := []byte("Alice")
	printDetailsWithReflectionByComparingType2(product7, customer7, payment7, true, 10, &number7, "Alice", 23.30, slice7)

	names := []string{"Alice", "Bob", "Charlie"}
	// Conversion du résultat en string
	val := selectValue(names, 1).(string)
	Printfln("Selected: %v", val)
}
