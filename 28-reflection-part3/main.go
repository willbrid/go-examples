package main

import (
	"fmt"
	"reflect"
	"strings"
)

/**
Les méthodes de type (reflect.Type) pour travailler avec des fonctions :
- NumIn() : cette méthode renvoie le nombre de paramètres définis par la fonction.
- In(index) : cette méthode retourne un Type qui reflète le paramètre à l'index spécifié.
- IsVariadic() : cette méthode renvoie vrai si le dernier paramètre est variadique.
- NumOut() : cette méthode renvoie le nombre de résultats définis par la fonction.
- Out(index) : cette méthode renvoie un Type qui reflète le résultat à l'index spécifié.


Les méthodes de type (reflect.Type) pour travailler avec des méthodes :
- NumMethod() : cette méthode renvoie le nombre de méthodes exportées définies pour le type de classe reflété.
- Method(index) : cette méthode renvoie la méthode reflétée à l'index spécifié, représenté par la classe Method.
- MethodByName(name) : cette méthode renvoie la méthode reflétée avec le nom spécifié. Les résultats sont une classe Method et
                       un booléen qui indique si une méthode portant le nom spécifié existe.

- Les champs définis par la classe Method :
--- Name : ce champ renvoie le nom de la méthode sous forme de chaîne.
--- PkgPath : ce champ est utilisé avec les interfaces, et non avec les méthodes accessibles via un type struct.
              Le champ renvoie une chaîne contenant le chemin du package. La chaîne vide est utilisée pour les champs exportés et
			  contiendra le nom du package de classe pour les champs non exportés.
--- Type : ce champ renvoie un Type (reflect.Type) qui décrit le type de fonction de la méthode.
--- Func : ce champ renvoie une valeur (reflect.Value) qui reflète la valeur de la fonction de la méthode.
           Lors de l'appel de la méthode, le premier argument doit être la classe sur laquelle la méthode est appelée.
--- Index : ce champ renvoie un int qui spécifie l'index de la méthode, à utiliser avec la méthode Method

- Les méthodes de valeur (reflect.Value) pour travailler avec des méthodes :
--- NumMethod() : cette méthode renvoie le nombre de méthodes exportées définies pour le type de structure reflété. Il appelle la méthode Type.NumMethod.
--- Method(index) : cette méthode renvoie une valeur (reflect.Value) qui reflète la fonction de la méthode à l'index spécifié.
                    Le récepteur n'est pas fourni comme premier argument lors de l'appel de la fonction.
--- MethodByName(name) : cette méthode renvoie une valeur (reflect.Value) qui reflète la fonction de la méthode avec le nom spécifié.
                         Le récepteur n'est pas fourni comme premier argument lors de l'appel de la fonction.


La méthode de valeur (reflect.Value) pour invoquer des fonctions :
- Call(params) : cette fonction invoque la fonction reflétée en utilisant la []reflect.Value comme paramètre.
                 Le résultat est une []reflect.Value qui contient les résultats de la fonction. Les valeurs fournies en paramètres doivent correspondre
				 à celles définies par la fonction.


Les fonctions du package reflect pour créer de nouveaux types de fonctions (reflect.Type) et valeurs de fonctions (reflect.Value) :
- FuncOf(params, results, variadic) : cette fonction crée un nouveau Type qui reflète un type de fonction avec les paramètres et les résultats spécifiés.
                                      Le dernier argument spécifie si le type de fonction a un paramètre variadique. Les paramètres et les résultats
									  sont spécifiés en tant que tranches de type.
- MakeFunc(type, fn) : cette fonction renvoie une valeur (reflect.Value) qui reflète une nouvelle fonction qui est un wrapper autour de la fonction fn.
                       La fonction doit accepter une tranche de valeur (reflect.Value) comme seul paramètre et renvoyer une tranche de valeur
					   comme seul résultat.
**/

func inspectFuncType(f interface{}) {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() == reflect.Func {
		Printfln("Function parameters : %v", funcType.NumIn())
		for i := 0; i < funcType.NumIn(); i++ {
			paramType := funcType.In(i)
			if i < (funcType.NumIn() - 1) {
				Printfln("Parameter #%v, Type : %v", i, paramType)
			} else {
				Printfln("Parameter #%v, Type : %v, Variadic : %v", i, paramType, funcType.IsVariadic())
			}
		}
		Printfln("Function results : %v", funcType.NumOut())
		for i := 0; i < funcType.NumOut(); i++ {
			resultType := funcType.Out(i)
			Printfln("Result #%v, Type : %v", i, resultType)
		}
	}
}

func invokeFunction(f interface{}, params ...interface{}) {
	paramVals := []reflect.Value{}
	for _, p := range params {
		paramVals = append(paramVals, reflect.ValueOf(p))
	}
	funcVal := reflect.ValueOf(f)
	if funcVal.Kind() == reflect.Func {
		results := funcVal.Call(paramVals)
		for i, r := range results {
			Printfln("Result #%v : %v", i, r)
		}
	}
}

func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}
	if sliceVal.Kind() == reflect.Slice &&
		mapperVal.Kind() == reflect.Func &&
		mapperVal.Type().NumIn() == 1 &&
		mapperVal.Type().In(0) == sliceVal.Type().Elem() {
		for i := 0; i < sliceVal.Len(); i++ {
			result := mapperVal.Call([]reflect.Value{sliceVal.Index(i)})
			for _, r := range result {
				mapped = append(mapped, r.Interface())
			}
		}
	}

	return
}

func mapSliceWithReflectFunction(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}

	if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func {
		paramTypes := []reflect.Type{sliceVal.Type().Elem()}
		resultTypes := []reflect.Type{}
		for i := 0; i < mapperVal.Type().NumOut(); i++ {
			resultTypes = append(resultTypes, mapperVal.Type().Out(i))
		}
		expectedFuncType := reflect.FuncOf(paramTypes, resultTypes, mapperVal.Type().IsVariadic())
		if mapperVal.Type() == expectedFuncType {
			for i := 0; i < sliceVal.Len(); i++ {
				result := mapperVal.Call([]reflect.Value{sliceVal.Index(i)})
				for _, r := range result {
					mapped = append(mapped, r.Interface())
				}
			}
		} else {
			Printfln("Function type not as expected")
		}
	}

	return
}

func makeMapperFunc(mapper interface{}) interface{} {
	mapVal := reflect.ValueOf(mapper)
	if mapVal.Kind() == reflect.Func && mapVal.Type().NumIn() == 1 && mapVal.Type().NumOut() == 1 {
		inType := reflect.SliceOf(mapVal.Type().In(0))
		inTypeSlice := []reflect.Type{inType}
		outType := reflect.SliceOf(mapVal.Type().Out(0))
		outTypeSlice := []reflect.Type{outType}
		funcType := reflect.FuncOf(inTypeSlice, outTypeSlice, false)
		funcVal := reflect.MakeFunc(funcType, func(params []reflect.Value) (results []reflect.Value) {
			srcSliceVal := params[0]
			resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), 10)
			for i := 0; i < srcSliceVal.Len(); i++ {
				r := mapVal.Call([]reflect.Value{srcSliceVal.Index(i)})
				resultsSliceVal.Index(i).Set(r[0])
			}
			results = []reflect.Value{resultsSliceVal}
			return
		})
		return funcVal.Interface()
	}
	Printfln("Unexpected types")
	return nil
}

func inspectMethods(s interface{}) {
	sType := reflect.TypeOf(s)
	if sType.Kind() == reflect.Struct || (sType.Kind() == reflect.Ptr && sType.Elem().Kind() == reflect.Struct) {
		Printfln("Type : %v, Methods : %v", sType, sType.NumMethod())
		for i := 0; i < sType.NumMethod(); i++ {
			method := sType.Method(i)
			Printfln("Method name : %v, Type : %v", method.Name, method.Type)
		}
	}
}

func main() {
	// Inspection d'une fonction
	inspectFuncType(Find)

	/**
	L'appel d'une fonction de cette manière n'est pas une exigence courante car le code appelant aurait pu simplement appeler
	la fonction directement, mais cet exemple rend l'utilisation de la méthode Call claire et souligne que les paramètres et
	les résultats sont tous deux exprimés à l'aide de tranches de valeur.
	**/
	names := []string{"Alice", "Bob", "Charlie"}
	invokeFunction(Find, names, "London", "Bob")

	results := mapSlice(names, strings.ToUpper)
	Printfln("Results : %v", results)

	name1s := []string{"Alice", "Bob", "Charlie"}
	result1s := mapSliceWithReflectFunction(name1s, strings.ToUpper)
	Printfln("Results 1 : %v", result1s)

	lowerStringMapper2 := makeMapperFunc(strings.ToLower).(func([]string) []string)
	name2s := []string{"Alice", "Bob", "Charlie"}
	result2s := lowerStringMapper2(name2s)
	Printfln("Lowercase Results : %v", result2s)

	incrementFloatMapper := makeMapperFunc(func(val float64) float64 {
		return val + 1
	}).(func([]float64) []float64)
	prices := []float64{279, 48.95, 19.50}
	floatResults := incrementFloatMapper(prices)
	Printfln("Increment Results : %v", floatResults)

	floatToStringMapper := makeMapperFunc(func(val float64) string {
		return fmt.Sprintf("$%.2f", val)
	}).(func([]float64) []string)
	Printfln("Price Results : %v", floatToStringMapper(prices))

	// Inspection d'une méthode d'une classe
	inspectMethods(Purchase{})
	inspectMethods(&Purchase{})
}
