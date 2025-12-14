package main

import "fmt"

func main() {
	/** Utilisation des opérateurs de comparaison **/

	/**
	On ne peut comparer que des valeurs du même type.
	Si les types sont différents, une conversion de type explicite est nécessaire avant la comparaison.

	== : cet opérateur renvoie vrai si les opérandes sont égaux.
	!= : cet opérateur renvoie vrai si les opérandes ne sont pas égaux.
	< : cet opérateur renvoie vrai si le premier opérande est inférieur au second opérande.
	> : cet opérateur renvoie vrai si le premier opérande est supérieur au second opérande.
	<= : cet opérateur renvoie vrai si le premier opérande est inférieur ou égal au deuxième opérande.
	>= : cet opérateur renvoie vrai si le premier opérande est supérieur ou égal au second opérande.
	**/
	const first int = 100
	const second int = 200
	var equal bool = first == second
	var notEqual bool = first != second
	var lessThan bool = first < second
	var lessThanOrEqual bool = first <= second
	var greaterThan bool = first > second
	var greaterThanOrEqual bool = first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)

	var max int
	if first > second {
		max = first
	} else {
		max = second
	}
	fmt.Println("MAX : ", max)

	// On peut comparer les pointeurs pour voir s'ils pointent vers le même emplacement mémoire.
	var alpha int = 100
	var beta int = 100
	var pointer1 *int = &alpha
	var pointer2 *int = &beta
	var pointer3 *int = &alpha
	fmt.Println("Comparaison des pointeurs")
	fmt.Println(pointer1 == pointer3)
	fmt.Println(pointer1 == pointer2)
	// Ces comparaisons suivent les pointeurs pour comparer les valeurs stockées aux emplacements mémoire référencés.
	fmt.Println("Comparaison des valeurs stockées par les pointeurs")
	fmt.Println(*pointer1 == *pointer3)
	fmt.Println(*pointer1 == *pointer2)

	/**
	Les opérateurs logiques comparent des valeurs booléennes. Les résultats produits par ces opérateurs peuvent être affectés à
	des variables ou utilisés dans une expression de contrôle de flux.

	|| : cet opérateur renvoie vrai si au moins un des opérandes est vrai. Si le premier opérande est vrai, le second ne sera pas évalué.
	&& : Cet opérateur renvoie vrai si les deux opérandes sont vrais. Si le premier opérande est faux, le second ne sera pas évalué.
	! : Cet opérateur s'utilise avec un seul opérande. Il renvoie vrai si l'opérande est faux et faux si l'opérande est vrai.
	**/
	var maxMph int = 50
	var passengerCapacity int = 4
	var airbags bool = true
	var familyCar bool = passengerCapacity > 2 && airbags
	var sportsCar = maxMph > 100 || passengerCapacity == 2
	var canCategorize = !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
}
