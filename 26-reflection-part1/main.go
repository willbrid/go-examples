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
- ConvertibleTo(type) : cette méthode renvoie true si le Type sur lequel la méthode est appelée peut être converti en Type spécifié.

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
- CanSet() : cette méthode renvoie true si la valeur peut être définie et false sinon.
- SetBool(val) : cette méthode définit la valeur sous-jacente sur le booléen spécifié.
- SetBytes(slice) : cette méthode définit la valeur sous-jacente sur la tranche d'octets spécifiée.
- SetFloat(val) : cette méthode définit la valeur sous-jacente sur le float64 spécifié.
- SetInt(val) : cette méthode définit la valeur sous-jacente sur l'int64 spécifié.
- SetUint(val) : cette méthode définit la valeur sous-jacente sur l'uint64 spécifié.
- SetString(val) : cette méthode définit la valeur sous-jacente sur la chaîne spécifiée.
- Set(val) : cette méthode définit la valeur sous-jacente sur la valeur sous-jacente de la valeur spécifiée.
- Convert(type) : cette méthode effectue une conversion de type et renvoie une valeur (reflect.Value) avec le nouveau type et la valeur d'origine.
- OverflowFloat(val) : cette méthode renvoie true si la valeur float64 spécifiée provoquerait un débordement si elle était convertie dans
                       le type de la valeur sur laquelle la méthode est appelée. Cette méthode paniquera à moins que la méthode Value.Kind
					   ne renvoie Float32 ou Float64.
- OverflowInt(val) : cette méthode renvoie true si la valeur int64 spécifiée provoquerait un débordement si elle était convertie dans
                     le type de la valeur sur laquelle la méthode est appelée. Cette méthode paniquera à moins que la méthode Value.Kind
					 ne renvoie l'un des types d'entiers signés.
- OverflowUint(val) : cette méthode renvoie true si la valeur uint64 spécifiée provoquerait un débordement si elle était convertie dans le type
                      de la valeur sur laquelle la méthode est appelée. Cette méthode paniquera à moins que la méthode Value.Kind ne renvoie
					  l'un des types d'entiers non signés.

La fonction de package reflect pour comparer les valeurs :
- DeepEqual(val, val) : cette fonction compare deux valeurs et renvoie true si elles sont identiques.
La fonction DeepEqual ne panique pas et effectue des comparaisons supplémentaires qui ne sont pas possibles avec l'opérateur ==.
Mais en général, la fonction DeepEqual effectue une comparaison en inspectant de manière récursive tous les champs ou éléments d'une valeur.
L'un des aspects les plus utiles de ce type de comparaison est que les tranches sont égales si toutes leurs valeurs sont égales, ce qui répond à l'une
des limitations les plus couramment rencontrées de l'opérateur de comparaison standard ==.

Les fonctions de création de nouvelles valeurs (reflect.Value) via le package reflect :
- New(type) : cette fonction crée une valeur qui pointe vers une valeur du type spécifié, initialisée à la valeur zéro du type.
              Il faut être prudent avec la fonction New car elle renvoie un pointeur vers une nouvelle valeur du type spécifié,
			  ce qui signifie qu'il est facile de créer un pointeur vers un pointeur.
- Zero(type) : cette fonction crée une valeur qui représente la valeur zéro du type spécifié.
- MakeMap(type) : cette fonction crée une nouvelle map
- MakeMapWithSize(type, size) : cette fonction crée une nouvelle map avec la taille spécifiée
- MakeSlice(type, capacity) : cette fonction crée une nouvelle tranche.
- MakeFunc(type, args, results) : cette fonction crée une nouvelle fonction avec les arguments et les résultats spécifiés.
- MakeChan(type, buffer) : cette fonction crée un nouveau canal avec la taille de buffer spécifiée.
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

func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}

func incrementOrUpperWithPointer(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.Kind() == reflect.Ptr {
			// Utilisation de la méthode reflect.Value.Elem() pour suivre le pointeur jusqu'à sa valeur
			elemValue = elemValue.Elem()
		}
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}

/*
*
La fonction setAll utilise une boucle for pour traiter son paramètre variadique et recherche des valeurs qui sont des pointeurs vers des valeurs
du même type que le paramètre src. Lorsqu'un pointeur correspondant est trouvé, la valeur à laquelle il fait référence est modifiée avec la méthode Set.
La majeure partie du code de la fonction setAll est chargée de vérifier que les valeurs sont compatibles et peuvent être définies, mais le résultat est
que l'utilisation d'une chaîne comme premier argument définit tous les arguments de chaîne suivants et l'utilisation d'un int définit toutes
les valeurs int suivantes.
*
*/
func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr && targetVal.Elem().Type() == srcVal.Type() && targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}

/*
*
La fonction containsWithPanic accepte une tranche et renvoie true si elle contient une valeur spécifiée. La tranche est énumérée à
l'aide des méthodes Len et Index.
Cette instruction applique l'opérateur de comparaison à la valeur à un index spécifique dans la tranche et à la valeur cible.
Mais, puisque la fonction containsWithPanic accepte tous les types, l'application paniquera si la fonction reçoit des types qui ne peuvent pas être comparés.
*
*/
func containsWithPanic(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				found = true
			}
		}
	}
	return
}

/*
*
La fonction containsWithoutPanic accepte une tranche et renvoie true si elle contient une valeur spécifiée. La tranche est énumérée à
l'aide des méthodes Len et Index. Une condition est ajoutée permettant à l'opérateur de comparaison de s'appliquer uniquement aux valeurs
dont les types sont comparables.
*
*/
func containsWithoutPanic(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	Printfln("Slice is it comparable ? %v", sliceVal.Type().Elem().Comparable())
	Printfln("Target is it comparable ? %v", targetType.Comparable())
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem().Comparable() && targetType.Comparable() {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				found = true
			}
		}
	}
	return
}

func containsWithDeepEqual(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
			}
		}
	}
	return
}

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}

	return
}

func IsInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func IsFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

/*
*
Cette fonction ajoute une protection contre les débordements lors de la conversion d'un type entier à un autre et
d'une valeur à virgule flottante à une autre.
*
*/
func convertImproved(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		if IsInt(targetVal) && IsInt(srcVal) && targetVal.OverflowInt(srcVal.Int()) {
			Printfln("Int overflow")
			return src, false
		} else if IsFloat(targetVal) && IsFloat(srcVal) && targetVal.OverflowFloat(srcVal.Float()) {
			Printfln("Float overflow")
			return src, false
		}
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}

	return
}

/*
*
Le Type passé à la fonction reflect.New est obtenu à partir du résultat reflect.Type.Elem pour l'une des valeurs de paramètre,
ce qui évite de créer un pointeur sur un pointeur. La méthode Set est utilisée pour définir la valeur temporaire et effectuer le swap.
*
*/
func swap(first interface{}, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstValue.Type() == secondValue.Type() && firstValue.Kind() == reflect.Ptr && firstValue.Elem().CanSet() && secondValue.Elem().CanSet() {
		temp := reflect.New(firstValue.Elem().Type())
		temp.Elem().Set(firstValue.Elem())
		firstValue.Elem().Set(secondValue.Elem())
		secondValue.Elem().Set(temp.Elem())
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

	/**
	Le résultat montrera qu'aucune des valeurs reçues par la fonction incrementOrUpper ne peut être définie.
	La méthode CanSet est source de confusion, mais n'oublions pas que les valeurs sont copiées lorsqu'elles sont utilisées comme arguments de fonctions
	et de méthodes. Lorsque les valeurs sont transmises à l'incrémentOrUpper, elles sont copiées.
	**/
	name1 := "Alice"
	price1 := 279
	city1 := "London"
	incrementOrUpper(name1, price1, city1)
	for _, val := range []interface{}{name1, price1, city1} {
		Printfln("Value: %v", val)
	}

	/**
	la réflexion ne peut modifier une valeur que si le stockage d'origine est accessible. Ici les pointeurs sont utilisés pour appeler la
	fonction incrementOrUpper, ce qui nécessite une modification du code de réflexion pour détecter les pointeurs et, lorsqu'il en trouve un,
	utiliser la méthode Elem pour suivre le pointeur jusqu'à sa valeur.
	**/
	name2 := "Alice"
	price2 := 279
	city2 := "London"
	incrementOrUpperWithPointer(&name2, &price2, &city2)
	for _, val := range []interface{}{name2, price2, city2} {
		Printfln("Value: %v", val)
	}

	name3 := "Alice"
	price3 := 279
	city3 := "London"
	setAll("New String", &name3, &price3, &city3)
	setAll(10, &name3, &price3, &city3)
	for _, val := range []interface{}{name3, price3, city3} {
		Printfln("Value: %v", val)
	}

	city4 := "London"
	citiesSlice4 := []string{"Paris", "Rome", "London"}
	sliceOfSlices4 := [][]string{citiesSlice4, {"First", "Second", "Third"}}
	Printfln("Found #1: %v", containsWithPanic(citiesSlice4, city4))
	//Printfln("Found #2: %v", containsWithPanic(sliceOfSlices4, citiesSlice4))
	Printfln("Found #2: %v", containsWithoutPanic(citiesSlice4, city4))
	Printfln("Found #3: %v", containsWithoutPanic(sliceOfSlices4, citiesSlice4))
	Printfln("Found #4: %v", containsWithDeepEqual(citiesSlice4, city4))
	Printfln("Found #5: %v", containsWithDeepEqual(sliceOfSlices4, citiesSlice4))

	/**
	Le premier appel à la fonction convert tente de convertir une valeur int en float64, ce qui réussit, et le deuxième appel tente de
	convertir une chaîne en float64, ce qui échoue.
	**/
	name5 := "Alice"
	price5 := 279
	newVal, ok := convert(price5, 100.00)
	Printfln("Converted %v : %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name5, 100.00)
	Printfln("Converted %v : %v, %T", ok, newVal, newVal)
	/**
	L'appel à la fonction convertImproved tente de convertir la valeur 5000 en un int8, ce qui provoquerait un débordement.
	La méthode OverflowInt renvoie true et la conversion n'est donc pas effectuée.
	**/
	price6 := 5000
	newVal, ok = convertImproved(price6, int8(100))
	Printfln("Converted %v : %v, %T", ok, newVal, newVal)

	name7 := "Alice"
	price7 := 279
	city7 := "London"
	swap(&name7, &city7)
	for _, val := range []interface{}{name7, price7, city7} {
		Printfln("Value : %v", val)
	}
}
