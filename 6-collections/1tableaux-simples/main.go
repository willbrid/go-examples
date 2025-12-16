package main

import "fmt"

func main() {
	/**
	Travailler avec des tableaux
	Les tableaux Go ont une longueur fixe et contiennent des éléments d'un seul type, accessibles par leur index.
	Les types de tableaux comprennent la taille du tableau entre crochets, suivie du type des éléments qu'il contient,
	appelé type sous-jacent. La longueur et le type des éléments d'un tableau sont fixes et sa longueur doit être spécifiée
	comme une constante.

	Le type d'un tableau est la combinaison de sa taille et de son type sous-jacent.
	Par exemple Le type de la variable « names » est [3]string, ce qui signifie un tableau dont le type sous-jacent est
	une chaîne de caractères et dont la capacité est de 3. Chaque combinaison de type sous-jacent et de capacité est un type distinct.
	Donc un tableau de [4]string est un type différent d'un tableau de [5]string, même si les deux contiennent des chaînes de caractères.
	**/
	var names [3]string
	fmt.Println(names)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	/**
	Les tableaux peuvent également être définies à l’aide d’une syntaxe littérale.
	Avec la syntaxe littérale, le type de tableau est suivi d'accolades contenant les éléments qui composeront le tableau.
	**/
	var names1 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names1)

	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names2)

	names3 := [5]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names3)

	var coords [3][3]int
	coords[0][1] = 7
	coords[1][2] = 10
	fmt.Println(coords)

	/**
	Par défaut, Go manipule les valeurs plutôt que les références. Ce comportement s'étend aux tableaux : assigner un tableau à
	une nouvelle variable permet de copier le tableau et ses valeurs.
	**/
	var names4 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names5 [3]string = names4
	names4[0] = "Canoe"
	fmt.Println("names4 : ", names4)
	fmt.Println("names5 : ", names5)

	/**
	Un pointeur peut être utilisé pour créer une référence à un tableau.
	Le type de la variable pointerNames6 est *[3]string, désignant un pointeur vers un tableau ayant la capacité de
	stocker trois valeurs de chaîne.
	**/
	var names6 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames6 *[3]string = &names6
	names6[0] = "Canoe"
	fmt.Println("Names : ", names6)
	fmt.Println("Pointer : ", *pointerNames6)

	var names7 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames7Element1 *string = &names7[1]
	fmt.Println("Pointer to Element 1 Before : ", *pointerNames7Element1)
	names7[1] = "Canoe"
	fmt.Println("Names : ", names7)
	fmt.Println("Pointer to Element 1 After : ", *pointerNames7Element1)

	// Les opérateurs de comparaison == et != peuvent être appliqués aux tableaux.
	// Deux tableaux sont égaux s'ils sont du même type et contiennent des éléments égaux dans le même ordre.
	var names8 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names9 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var isSame bool = names8 == names9
	fmt.Println("Comparaison : ", isSame)

	// Les tableaux sont énumérés à l'aide des mots-clés `for` et `range`.
	var names10 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names10 {
		fmt.Println("Index : ", index, " - Value : ", value)
	}

	var names11 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, value := range names11 {
		fmt.Println("Value : ", value)
	}
}
