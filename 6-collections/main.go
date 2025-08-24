package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections")

	/** Travailler avec des tableaux **/
	fmt.Println("Travailler avec des tableaux")
	var names [3]string
	fmt.Println(names)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	// Les tableaux peuvent également être définies à l’aide d’une syntaxe littérale.
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

	var names4 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names5 [3]string = names4

	names4[0] = "Canoe"
	fmt.Println("names4 : ", names4)
	fmt.Println("names5 : ", names5)

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

	var names8 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names9 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var isSame bool = names8 == names9
	fmt.Println("Comparaison : ", isSame)

	var names10 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names10 {
		fmt.Println("Index : ", index, " - Value : ", value)
	}

	var names11 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, value := range names11 {
		fmt.Println("Value : ", value)
	}

	/** Travailler avec Tranches : Tableau dont on ne connait pas sa longueur en avance ou dont sa longueur est variable.
		Les tranches en Go sont des tableaux de taille variable, pratiques quand le nombre d’éléments est inconnu ou changeant.
		On peut les créer avec la fonction make : make(Type slice, longueur, capacité).
		La fonction make permet d'initialiser une tranche lors de sa création.
		Une tranche en Go est basée sur un tableau sous-jacent et contient trois informations :
		- un pointeur vers ce tableau,
		- sa longueur (nombre d’éléments qu'elle peut stocker (nombre d’éléments accessibles))
		- sa capacité (nombre d’éléments qui peut être stocké dans le tableau sous-jacent).

		La fonction len permet de déterminer la longueur d'une tranche : len(slicename)
		La fonction cap permet de déterminer la capacité d'une tranche : cap(slicename)
		La capacité sera toujours au moins égale à la longueur, mais peut être supérieure si une
		capacité supplémentaire a été allouée via la fonction make

		Avantage des tranches en Go : elles peuvent être agrandies avec la fonction append, qui ajoute de nouveaux éléments en créant
		si besoin un tableau plus grand, en copiant l’existant et en y ajoutant les nouvelles valeurs.
		Le résultat de la fonction append est une tranche dont la longueur a augmenté mais qui est toujours soutenue par
		le même tableau sous-jacent.
		La fonction append peut être utilisée pour ajouter une tranche à une autre.
	**/
	fmt.Println("Travailler avec Tranches : Tableau dont on ne connait pas sa longueur en avance ou dont sa longueur est variable.")
	var names12 []string = make([]string, 3)
	names12[0] = "Kayak"
	names12[1] = "Lifejacket"
	names12[2] = "Paddle"
	fmt.Println(names12)

	// Les tranches peuvent également être définies à l’aide d’une syntaxe littérale.
	names13 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names13)

	names14 := []string{"Kayak", "Lifejacket", "Paddle"}
	names14 = append(names14, "Hat", "Gloves")
	fmt.Println(names14)

	var names15 []string = []string{"Kayak", "Lifejacket", "Paddle"}
	var appendedNames15 []string = append(names15, "Hat", "Gloves")
	names15[0] = "Canoe"
	fmt.Println(names15)
	fmt.Println(appendedNames15)

	var names16 []string = make([]string, 3, 7)
	names16[0] = "Kayak"
	names16[1] = "Lifejacket"
	names16[2] = "Paddle"
	fmt.Println("Tableau : ", names16)
	fmt.Println("Longueur : ", len(names16))
	fmt.Println("Capacité : ", cap(names16))

	var names17 []string = make([]string, 3, 7)
	names17[0] = "Kayak"
	names17[1] = "Lifejacket"
	names17[2] = "Paddle"
	moreName := []string{"Hat Gloves"}
	appendedNames17 := append(names17, moreName...)
	fmt.Println("AppendedNames : ", appendedNames17)

	/**
		Créer une tranche à partir d'un tableau.
		- par défaut une tranche créée est basée sur un tableau sous-jacent.
		- si une tranche est créée à partir d'un tableau alors ce tableau devient le tableau sous-jacent de la tranche.
		- si on a une tranche avec la plage [x:y] créée à partir d'un tableau, sa capacité = (y-x) + 1
		- toutes les tranches créées à partir d'un tableau, partagent le même tableau.
		- Lorsqu’une tranche créée à partir d'un tableau existant est étendue avec append, elle utilise d’abord l’espace libre
		  de ce tableau existant. Si la capacité est atteinte, un nouveau tableau est créé, les données sont copiées,
		  et la tranche pointe alors vers ce nouveau tableau.

		Les plages sont exprimées entre crochets, les valeurs minimales et maximales étant séparées par deux points.
		Le premier index de la tranche correspond à la valeur minimale, et la longueur est le résultat de la soustraction de la
		valeur maximale et de la valeur minimale. Cela signifie que la plage [1:3] crée une plage dont l'index zéro est mappé à
		l'index 1 du tableau et dont la longueur est 2.
		L'index de début et le nombre peuvent être omis d'une plage pour inclure tous les éléments de la source.

		RENDRE LES TRANCHES PRÉVISIBLES : il faut traiter deux catégories de tranches en Go
	    - Vue fixe d’un tableau fixe → on peut modifier les éléments mais pas en ajouter, et les changements affectent toutes les tranches liées.
	    - Collection à taille variable → chaque tranche a son propre tableau, ce qui permet d’ajouter des éléments sans impacter d’autres tranches.

		- Spécifier de la capacité lors de la création d'une tranche à partir d'un tableau
		Les plages peuvent inclure une capacité maximale, ce qui permet de contrôler le moment où les tableaux seront dupliquées.
		La valeur maximale ne spécifie pas directement la capacité maximale. Celle-ci est déterminée en soustrayant la valeur minimale de
		la valeur maximale. Dans l'exemple avec la plage [1:3:3], la valeur maximale est de 3 et la valeur minimale de 1, ce qui signifie
		que la capacité sera limitée à 2. Par conséquent, l'opération d'ajout entraîne le redimensionnement de la tranche et l'allocation
		de son propre tableau, au lieu de l'étendre dans le tableau existant.
		si on a une tranche avec la plage [x:y:z] créée à partir d'un tableau, sa capacité = (z-x) et y est sa longueur
		**/
	var products [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Gloves")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Boots")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))

	var products1 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames1 := products1[1:3:3]
	allNames1 := products1[:]
	someNames1 = append(someNames1, "Gloves")
	fmt.Println("SomeNames : ", someNames1)
	fmt.Println("SomeNames len : ", len(someNames1), " cap : ", cap(someNames1))
	fmt.Println("AllNames : ", allNames1)
	fmt.Println("AllNames len : ", len(allNames1), " cap : ", cap(allNames1))

	var products2 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames2 := products2[1:]
	someNames2 := make([]string, 2)
	copy(someNames2, allNames2)
	fmt.Println("SomeNames2 : ", someNames2)
	fmt.Println("AllNames2 : ", allNames2)

	var products3 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames3 := products3[1:]
	var someNames3 []string // Variable non initialisée : la copie n'est pas effective
	copy(someNames3, allNames3)
	fmt.Println("SomeNames3 : ", someNames3)
	fmt.Println("AllNames3 : ", allNames3)

	/**
	La fonction copy en Go permet de copier des éléments d’une tranche vers une autre, garantissant des tableaux distincts.
	Elle peut dupliquer tout ou partie d’une tranche. La copie s’arrête à la fin de la tranche source ou cible,
	mais ne redimensionne pas la tranche cible, qui doit donc avoir une longueur suffisante.

	Si la tranche de destination n’est pas initialisée, sa longueur est nulle et copy ne copie aucun élément.
	Aucune erreur n’est générée, mais cela conduit souvent à une tranche vide inattendue.
	**/
	var products4 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames4 := products4[1:]
	someNames4 := []string{"Boots", "Canoe"}
	copy(someNames4[1:], allNames4[2:3]) // Le tableau source allNames4 sera copié à partir de la position 2 et
	// La copie va être positionné à partir de la position 1 du tableau destination someNames4
	fmt.Println("SomeNames4 : ", someNames4)
	fmt.Println("AllNames4 : ", allNames4)

	/**
	On peut spécifier des plages lors de la copie de tranches.
	Si la tranche de destination est plus grande que la tranche source, la copie se poursuit jusqu'à la copie du dernier élément de la source.
	La tranche source ne contient que deux éléments et aucune plage n'est utilisée. Par conséquent, la fonction de copie commence à copier
	les éléments de la tranche replacementProducts vers la tranche products5 et s'arrête lorsque la fin de la tranche replacementProducts
	est atteinte. Les éléments restants de la tranche products ne sont pas affectés par la copie.
	**/
	var products5 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products5[:], replacementProducts)
	fmt.Println("products5 : ", products5) // La copie sera effective uniquement sur les deux premiers éléments du tableau products5

	/**
	Si la tranche de destination est plus petite que la tranche source, la copie continue jusqu'à ce que tous les éléments
	de la tranche de destination aient été remplacés
	**/
	var products6 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts1 := []string{"Canoe", "Boots"}
	copy(products6[0:1], replacementProducts1)
	fmt.Println("products6 : ", products6) // La copie sera effective uniquement sur le premier élément du tableau products6

	var products7 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products7[:2], products7[3:]...)
	fmt.Println("Deleted : ", deleted) // On forme le tableau deleted en ajoutant au deux premiers éléments (0 et 1),
	// l'élément numéro 3 : d'où cela supprime l'élément numéro 2

	var products8 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products8[2:] {
		fmt.Println("Index : ", index, " - Value: ", value)
	}

	var products9 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	sort.Strings(products9[:])
	for index, value := range products9 {
		fmt.Println("Index : ", index, " - Value: ", value)
	}

	products10 := products9
	fmt.Println("Equal : ", reflect.DeepEqual(products9, products10)) // Comparaison de tableau avec la fonction DeepEqual du package reflect

	/**
	    Obtenir le tableau sous-jacent à une tranche
	    Si nous avons une tranche mais que nous avez besoin d'un tableau, généralement parce qu'une fonction en requiert un comme argument,
		nous pouvons effectuer une conversion explicite sur la tranche. Nous devons être prudent lors de la spécification du type de tableau,
		car une erreur se produira si le nombre d'éléments requis par le tableau dépasse la longueur de la tranche. La longueur du tableau
		peut être inférieure à celle de la tranche, auquel cas le tableau ne contiendra pas toutes les valeurs de la tranche.
		**/
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println("Array from slice p1 :", array)

	/** Travailler avec Maps : tableau associatif (dictionnaire) clé-valeur
	Les dictionnaires sont une structure de données intégrée qui associe des valeurs de données à des clés. Ils sont créées avec la
	fonction intégrée make, comme pour les tranches. Le type d'un dictionnaire est spécifié à l'aide du mot-clé map,
	suivi du type de clé entre crochets, puis du type de valeur.
	**/
	fmt.Println("Travailler avec Maps : tableau associatif clé-valeur")
	var products11 map[string]float64 = make(map[string]float64, 10) // string représente le type de la clé et float64 le type de la valeur
	products11["Kayak"] = 279
	products11["Lifejacket"] = 48.95
	fmt.Println("Map size: ", len(products11))
	fmt.Println("Price: ", products11["Kayak"])
	fmt.Println("Price: ", products11["Hat"])

	// Les dictionnaires peuvent également être définies à l’aide d’une syntaxe littérale.
	var products12 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size: ", len(products12))
	fmt.Println("Price: ", products12["Kayak"])
	fmt.Println("Price: ", products12["Hat"])

	/**
	Les dictionnaires renvoient le zéro pour le type de valeur lors de la lecture d'une clé inexistante. Il peut alors être difficile
	de distinguer une valeur stockée avec le zéro du type de valeur d'une clé inexistante. Pour résoudre ce problème,
	les dictionnaires produisent deux valeurs lors de la lecture d'une valeur :
	--- la première valeur est soit la valeur associée à la clé spécifiée, soit le zéro du type de valeur en l'absence de clé.
	--- la seconde valeur est une valeur booléenne qui est vraie si le dictionnaire contient la clé spécifiée et fausse dans
	    le cas contraire. Cette seconde valeur est généralement affectée à une variable nommée ok, d'où le terme « comma ok ».
	**/
	var products13 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	value, ok := products13["Hat"]
	if ok {
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}
	delete(products13, "Hat") // Supprimer un élément du tableau associatif à partir de sa clé
	if value1, ok1 := products13["Hat"]; ok1 {
		fmt.Println("Stored value: ", value1)
	} else {
		fmt.Println("No stored value")
	}

	for key, value := range products13 {
		fmt.Println("Key : ", key, " - Value: ", value)
	}

	/**
	Il n'y a aucune garantie que le contenu d'un dictionnaire soit énuméré dans un ordre spécifique. Pour classer les valeurs d'un
	dictionnaire, la meilleure approche consiste à énumérer le dictionnaire, à créer une tranche contenant les clés,
	à trier la tranche, puis à énumérer la tranche pour lire les valeurs du dictionnaire.
	**/
	var products14 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	var keys []string = make([]string, 0, len(products14))
	for key := range products14 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key : ", key, " - Valeur : ", products14[key])
	}

	/** Comprendre la double nature des chaines de caractères
	Go traite les chaînes de caractères comme des tableaux d'octets et prend en charge la notation d'index de tableau et de plage de tranches.
	--- Lorsque la notation d'index est utilisée, le résultat est un octet de l'emplacement spécifié dans la chaîne
	--- Lorsqu'une chaîne est découpée en tranches, la tranche est également décrite à l'aide d'octets, mais le résultat est une chaîne
	**/
	fmt.Println("Comprendre la double nature des chaines de caractères")
	var price string = "$48.95"
	var currencyByte byte = price[0] // Byte est l'alias à uint8
	var currencyString string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency Byte : ", currencyByte)
	fmt.Println("Currency String : ", currencyString)
	fmt.Println("Length : ", len(price))
	if parseErr == nil {
		fmt.Println("Amount : ", amount)
	} else {
		fmt.Println("Parse Error : ", parseErr)
	}

	/**
	Le type rune représente un point de code Unicode, qui est essentiellement un caractère unique. Pour éviter de découper
	les chaînes au milieu des caractères, une conversion explicite en tranche de rune peut être effectuée.
	**/
	var price1 []rune = []rune("€48.95") // rune est l'alias à int32
	var currency1 string = string(price1[0])
	var amountString1 string = string(price1[1:])
	amount1, parseErr1 := strconv.ParseFloat(amountString1, 64)
	fmt.Println("Currency String : ", currency1)
	fmt.Println("Length : ", len(price1))
	if parseErr1 == nil {
		fmt.Println("Amount : ", amount1)
	} else {
		fmt.Println("Parse Error : ", parseErr1)
	}

	var price2 string = "$48.95"
	for index, char := range price2 {
		// Value correspondance en byte et string(char) correspondance en caractère
		fmt.Println("Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Notez que les valeurs d'index ne sont pas séquentielles. La boucle for traite la chaîne comme une séquence de caractères dérivée
	de la séquence d'octets sous-jacente. Les valeurs d'index correspondent au premier octet composant chaque caractère.
	La deuxième valeur d'index est 3, car le premier caractère de la chaîne est composé d'octets en positions 0, 1 et 2.
	**/
	var price3 string = "€48.95"
	for index, char := range price3 {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price3 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Pour énumérer les octets sous-jacents sans les convertir en caractères, nous pouvons effectuer une conversion explicite
	en tranche d'octets.
	**/
	var price4 string = "€48.95"
	for index, char := range []byte(price4) {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price4 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Pour énumérer les caractères de la chaine, nous pouvons aussi effectuer une conversion explicite en tranche de rune.
	**/
	var price5 []rune = []rune("€48.95")
	for index, char := range price5 {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price5 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}
}
