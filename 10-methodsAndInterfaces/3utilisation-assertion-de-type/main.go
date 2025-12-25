package main

import "fmt"

/**
Les assertions de type ne s'appliquent qu'aux interfaces et servent à indiquer au compilateur qu'une valeur d'interface possède
un type dynamique spécifique. Les conversions de type, quant à elles, ne s'appliquent qu'à des types spécifiques, et non aux interfaces,
et seulement si la structure de ces types est compatible, par exemple entre des types struct ayant les mêmes champs.
**/

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	/**
	Il est souvent utile de pouvoir accéder directement au type dynamique, ce que l'on appelle le rétrécissement de type :
	le processus qui consiste à passer d'un type moins précis à un type plus précis.
	Une assertion de type permet d'accéder au type dynamique d'une valeur d'interface.

	Une assertion de type est effectuée en appliquant un point après une valeur d'interface, suivi du type cible entre parenthèses.
	Une fois appliquée, l'on peut utiliser tous les champs et méthodes définis pour ce type cible, et pas seulement ceux
	définis par l'interface.
	**/

	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
	}
	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service :", s.description, "Price :", s.monthlyFee*float64(s.durationMonths))
	}

	/**
	L'on peut rencontrer des problèmes si expenses contient autres valeurs de structure qu'une valeur de Service. Par exemple si expenses
	contient les valeurs de Service et de Product alors le code ci-après donnera une erreur fatale :
	-> s := expense.(Service)
	   fmt.Println("Service :", s.description, "Price :", s.monthlyFee*float64(s.durationMonths))
	Comme solution, il existe une forme particulière d'assertion de type qui indique si une assertion peut être effectuée.
	Les assertions de type peuvent produire deux résultats. Le premier résultat est l'attribution du type dynamique, et le second
	résultat est un booléen indiquant si l'assertion a pu être exécutée.
	**/
	otherExpenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range otherExpenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service :", s.description, "Price :", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense :", expense.getName(), "Cost :", expense.getCost(true))
		}
	}

	/**
	Les instructions switch de Go peuvent être utilisées pour accéder aux types dynamiques, ce qui peut constituer une manière
	plus concise d'effectuer des assertions de type avec des instructions if.

	L'instruction switch utilise une assertion de type spéciale qui utilise le mot-clé `type`.
	Chaque instruction `case` spécifie un type et un bloc de code qui sera exécuté lorsque la valeur évaluée par l'instruction `switch`
	est du type spécifié. Le compilateur Go est capable de comprendre la relation entre les valeurs évaluées par l'instruction `switch`
	et n'autorise pas les instructions `case` pour des types incompatibles.
	Par exemple, le compilateur générera une erreur si une instruction `case` est présente pour le type `Product`,
	car l'instruction `switch` évalue des valeurs de type `Expense` et le type `Product` ne possède pas les méthodes nécessaires
	pour implémenter l'interface puisque les méthodes du fichier `product.go` utilisent des récepteurs de pointeurs.

	Une instruction par défaut peut être utilisée pour spécifier un bloc de code qui sera exécuté lorsqu'aucune des instructions
	« case » ne correspond.
	**/

	someOtherExpenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range someOtherExpenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service :", value.description, "Price :", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product :", value.name, "Price :", value.price)
		default:
			fmt.Println("Expense :", expense.getName(), "Cost :", expense.getCost(true))
		}
	}
}
