package main

import (
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


La méthode de valeur (reflect.Value) pour invoquer des fonctions :
- Call(params) : cette fonction invoque la fonction reflétée en utilisant la []reflect.Value comme paramètre.
                 Le résultat est une []reflect.Value qui contient les résultats de la fonction. Les valeurs fournies en paramètres doivent correspondre
				 à celles définies par la fonction.

La fonction du package reflect pour créer de nouveaux types de fonctions (reflect.Type) et valeurs de fonctions (reflect.Value) :
- FuncOf(params, results, variadic) : cette fonction crée un nouveau Type qui reflète un type de fonction avec les paramètres et les résultats spécifiés.
                                      Le dernier argument spécifie si le type de fonction a un paramètre variadique. Les paramètres et les résultats
									  sont spécifiés en tant que tranches de type.
- MakeFunc(type, fn) : cette fonction renvoie une valeur qui reflète une nouvelle fonction qui est un wrapper autour de la fonction fn.
                       La fonction doit accepter une tranche de valeur comme seul paramètre et renvoyer une tranche de valeur comme seul résultat.
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
}
