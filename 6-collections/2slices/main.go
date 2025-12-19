package main

import "fmt"

func main() {
	/**
	La meilleure façon de concevoir les tranches est de les considérer comme un tableau de longueur variable, car elles sont utiles
	lorsque nous ne connaissons pas le nombre de valeurs à stocker ou lorsque ce nombre évolue dans le temps.
	Une façon de définir une tranche est d'utiliser la fonction intégrée `make`. La fonction make accepte des arguments qui spécifient le type
	et la longueur de la tranche.
	**/

	names := make([]string, 3)
	names[0] = "kayak"
	names[1] = "lifejacket"
	names[2] = "paddle"
	fmt.Println("names :", names)

	// Syntaxe littérale de tranche
	name1s := []string{"kayak", "lifejacket", "paddle"}
	fmt.Println("name1s :", name1s)

	/**
	La combinaison du type et de la longueur d'une tranche permet de créer un tableau, qui sert de zone de stockage pour cette tranche.
	Une tranche est une structure de données qui contient trois valeurs : un pointeur vers le tableau, sa longueur et sa capacité.
	La longueur d'une tranche correspond au nombre d'éléments qu'elle peut stocker, et sa capacité correspond au nombre total
	d'éléments pouvant être stockés dans le tableau.
	Les tranches prennent en charge la notation d'index de type tableau, ce qui permet d'accéder aux éléments du tableau sous-jacent.
	**/

	/**
	La fonction intégrée `append` accepte une tranche et un ou plusieurs éléments à ajouter à la tranche, séparés par des virgules.
	La fonction `append` crée un tableau suffisamment grand pour contenir les nouveaux éléments, copie le tableau existant et
	y ajoute les nouvelles valeurs. Le résultat de la fonction `append` est une tranche qui est mappée sur le nouveau tableau.
	**/
	name2s := []string{"kayak", "lifejacket", "paddle"}
	name2s = append(name2s, "Hat", "Gloves")
	fmt.Println("name2s :", name2s)

	// La tranche originale et son tableau sous-jacent existent toujours et peuvent être utilisés.
	name3s := []string{"kayak", "lifejacket", "paddle"}
	appendedNames := append(name3s, "Hat", "Gloves")
	name3s[0] = "Canoe"
	fmt.Println("name3s :", name3s)
	fmt.Println("appendedNames :", appendedNames)

	/**
	Comme précisé plus haut, Les tranches ont une longueur et une capacité. La longueur d'une tranche correspond au nombre de valeurs
	qu'elle peut contenir, tandis que la capacité représente le nombre d'éléments pouvant être stockés dans le tableau sous-jacent
	avant que la tranche ne doive être redimensionnée et un nouveau tableau créé. La capacité est toujours au moins égale à la longueur,
	mais peut être supérieure si de la capacité supplémentaire a été allouée avec la fonction `make`.

	Les fonctions intégrées `len` et `cap` renvoient la longueur et la capacité d'une tranche.

	Si nous définissons une variable de type tranche sans l'initialiser, la tranche résultante aura une longueur et une capacité nulles,
	ce qui provoquera une erreur lors de l'ajout d'un élément.
	**/
	name4s := make([]string, 3, 6)
	name4s[0] = "kayak"
	name4s[1] = "lifejacket"
	name4s[2] = "paddle"
	fmt.Println("name4s len :", len(name4s))
	fmt.Println("name4s cap :", cap(name4s))

	/**
	Le tableau sous-jacent n'est pas remplacé lorsque la fonction `append` est appelée sur une tranche ayant une capacité suffisante
	pour accueillir les nouveaux éléments.
	Dans cet exemple, la fonction `append` a pour résultat une tranche dont la longueur a augmenté, mais qui repose toujours sur
	le même tableau sous-jacent. La tranche d'origine existe toujours et repose sur le même tableau, ce qui a pour effet
	qu'il existe désormais deux vues d'un même tableau.
	**/
	name5s := make([]string, 3, 6)
	name5s[0] = "kayak"
	name5s[1] = "lifejacket"
	name5s[2] = "paddle"
	appendedName1s := append(name5s, "Hat", "Gloves")
	name5s[0] = "Canoe"
	fmt.Println("name5s :", name5s)
	fmt.Println("appendedName1s :", appendedName1s)

	/**
	La fonction `append` peut être utilisée pour ajouter une tranche à une autre.
	Le deuxième argument est suivi de trois points de suspension (`...`), ce qui est nécessaire car la fonction `append`
	intégrée définit un paramètre variadique.
	**/
	name6s := make([]string, 3, 6)
	name6s[0] = "kayak"
	name6s[1] = "lifejacket"
	name6s[2] = "paddle"
	moreNames := []string{"Hat", "Gloves"}
	appendedName2s := append(name6s, moreNames...)
	fmt.Println("appendedName2s :", appendedName2s)

	/**
	Il est possible de créer des tranches à partir de tableaux existants.
	La variable « products » se voit attribuer un tableau standard de longueur fixe contenant des chaînes de caractères.
	Ce tableau sert à créer des tranches à l'aide d'une plage, qui spécifie les valeurs minimale et maximale.

	Les intervalles sont exprimés entre crochets, les valeurs minimale et maximale étant séparées par un deux-points.
	Le premier indice de l'intervalle correspond à la valeur minimale, et sa longueur est égale à la
	différence entre la valeur maximale et la valeur minimale.
	Ainsi, l'intervalle [1:3] définit un intervalle dont l'indice zéro correspond à l'indice 1 du tableau et dont la longueur est de 2.
	L'indice de départ et le nombre d'éléments peuvent être omis d'une plage pour inclure tous les éléments de la source.

	Le code de l'exemple ci-dessous crée deux tranches, toutes deux basées sur le même tableau sous-jacent.
	La relation entre la tranche et le tableau existant peut engendrer des résultats différents lors de l'ajout d'éléments.

	Comme l'illustre l'exemple ci-dessous, il est possible de décaler une tranche afin que son premier indice ne corresponde pas au début
	du tableau et que son dernier indice ne pointe pas vers le dernier élément du tableau. Ainsi, l'indice 0 de la tranche `someProducts`
	correspond à l'indice 1 du tableau. Jusqu'à présent, la capacité des tranches était alignée sur la longueur du tableau sous-jacent,
	mais ce n'est plus le cas, car le décalage réduit la portion du tableau utilisable par la tranche.
	**/
	products := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	someProducts := products[1:3]
	allProducts := products[:]
	fmt.Println("someProducts :", someProducts)
	fmt.Println("someProducts len :", len(someProducts), " cap :", cap(someProducts))
	fmt.Println("allProducts :", allProducts)
	fmt.Println("allProducts len :", len(allProducts), " cap :", cap(allProducts))

	/**
		Cette tranche peut accueillir le nouvel élément sans être redimensionnée, mais l'emplacement du tableau qui servira à stocker cet élément
		est déjà inclus dans la tranche `allProduct1s`. Par conséquent, l'opération d'ajout agrandit la tranche `someProduct1s`
		et modifie l'une des valeurs accessibles via la tranche `allProduct1s`.

		Si nous utilisons une tranche comme vue fixe d'un tableau, nous pouvons nous attendre à ce que plusieurs tranches nous donnent une
		vue cohérente de ce tableau, et toutes les nouvelles valeurs que nous attribuons seront reflétées par toutes les tranches qui
		correspondent à l'élément modifié.

		L'ajout de la valeur 'Gloves' à la tranche someProduct1s modifie la valeur renvoyée par allProduct1s[3] car
		les tranches partagent le même tableau sous-jacent.
	    Le résultat montre également que la longueur et la capacité des tranches sont identiques, ce qui signifie qu'il n'est plus
		possible d'agrandir la tranche sans créer un tableau sous-jacent plus grand.
	**/
	product1s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	someProduct1s := product1s[1:3]
	allProduct1s := product1s[:]
	someProduct1s = append(someProduct1s, "Gloves")
	fmt.Println("someProduct1s :", someProduct1s)
	fmt.Println("someProduct1s len :", len(someProduct1s), " cap :", cap(someProduct1s))
	fmt.Println("allProduct1s :", allProduct1s)
	fmt.Println("allProduct1s len :", len(allProduct1s), " cap :", cap(allProduct1s))

	/**
	Le premier appel à la fonction `append` étend la tranche `someProduct2s` au sein du tableau sous-jacent existant.
	L'espace disponible étant limité lors du second appel à `append`, un nouveau tableau est créé, son contenu est copié et
	les deux tranches sont désormais associées à des tableaux différents.

	Le processus de redimensionnement ne copie que les éléments du tableau qui sont mappés par la tranche, ce qui a pour effet
	de réaligner les indices de la tranche et du tableau.
	**/
	product2s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	someProduct2s := product2s[1:3]
	allProduct2s := product2s[:]
	someProduct2s = append(someProduct2s, "Gloves")
	someProduct2s = append(someProduct2s, "Boots")
	fmt.Println("someProduct2s :", someProduct2s)
	fmt.Println("someProduct2s len :", len(someProduct2s), " cap :", cap(someProduct2s))
	fmt.Println("allProduct2s :", allProduct2s)
	fmt.Println("allProduct2s len :", len(allProduct2s), " cap :", cap(allProduct2s))

	/**
		Spécification de la capacité lors de la création d'une tranche à partir d'un tableau

	    Les plages peuvent inclure une capacité maximale, ce qui permet de contrôler, dans une certaine mesure, le moment où les
		tableaux seront dupliqués.
		La valeur maximale ne spécifie pas directement la capacité maximale. Celle-ci est déterminée en soustrayant la valeur minimale
		de la valeur maximale. Dans l'exemple ci-dessous, la valeur maximale est 3 et la valeur minimale est 1, ce qui limite la capacité à 2.
		Par conséquent, l'opération d'ajout entraîne le redimensionnement de la tranche et l'allocation de son propre tableau,
		au lieu de son extension dans le tableau existant. Le redimensionnement de la tranche signifie que la valeur 'Gloves' ajoutée
		à la tranche someProduct3s ne fait pas partie des valeurs mappées par la tranche allProduct3s.
	**/
	product3s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	someProduct3s := product3s[1:3:3]
	allProduct3s := product3s[:]
	someProduct3s = append(someProduct3s, "Gloves")
	fmt.Println("someProduct3s :", someProduct3s)
	fmt.Println("someProduct3s len :", len(someProduct3s), " cap :", cap(someProduct3s))
	fmt.Println("allProduct3s :", allProduct3s)
	fmt.Println("allProduct3s len :", len(allProduct3s), " cap :", cap(allProduct3s))

	/**
	Création de tranches à partir d'autres tranches
	Il est également possible de créer des tranches à partir d'autres tranches, mais la relation entre les tranches n'est pas conservée
	si elles sont redimensionnées.

	Les tranches sont essentiellement des pointeurs vers des sections de tableaux, ce qui signifie qu'elles ne peuvent pas pointer
	vers une autre tranche. En réalité, les plages servent à déterminer les correspondances pour les tranches qui sont sous-jacentes
	par le même tableau.
	**/
	product4s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	allProduct4s := product4s[1:]
	someProduct4s := allProduct4s[1:3]
	allProduct4s = append(allProduct4s, "Gloves")
	allProduct4s[1] = "Canoe"
	fmt.Println("someProduct4s :", someProduct4s)
	fmt.Println("allProduct4s :", allProduct4s)

	/**
	La fonction `copy` permet de copier des éléments entre les tranches. Elle peut être utilisée pour garantir que les tranches
	possèdent des tableaux distincts et pour créer des tranches combinant des éléments provenant de différentes sources.
	**/
	/**
	Utilisation de la fonction `copy` pour garantir la séparation des tableaux de tranches
	La fonction `copy` permet de dupliquer une tranche existante, en sélectionnant tout ou partie de ses éléments, tout en veillant
	à ce que la nouvelle tranche soit associée à son propre tableau.
	La fonction `copy` accepte deux arguments : la tranche de destination et la tranche source.

	Cette fonction `copy` les éléments vers la tranche cible. Les tranches n'ont pas besoin d'avoir la même longueur, car
	la fonction `copy` s'arrête à la fin de la tranche source ou de destination. La tranche de destination n'est pas redimensionnée,
	même si le tableau sous-jacent dispose d'espace libre ; nous devons donc nous assurer qu'elle est suffisamment longue pour contenir
	le nombre d'éléments à copier.
	**/
	product5s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	allProduct5s := product5s[1:]
	someProduct5s := make([]string, 2)
	copy(someProduct5s, allProduct5s)
	fmt.Println("someProduct5s :", someProduct5s)
	fmt.Println("allProduct5s :", allProduct5s)

	/**
	Une erreur fréquente consiste à tenter de copier des éléments dans une tranche non initialisée.
	Dans l'exemple ci-dessous, aucun élément n'a été copié dans la tranche de destination. Cela se produit car les tranches non initialisées
	ont une longueur et une capacité nulles. La fonction `copy` s'arrête lorsque la longueur de la tranche de destination est atteinte,
	et comme cette longueur est nulle, aucune copie n'a lieu. Aucune erreur n'est signalée car la fonction `copy` a fonctionné comme prévu.
	**/
	product6s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	allProduct6s := product6s[1:]
	var someProduct6s []string
	copy(someProduct6s, allProduct6s)
	fmt.Println("someProduct6s :", someProduct6s)
	fmt.Println("allProduct6s :", allProduct6s)

	/**
	Un contrôle précis des éléments copiés peut être obtenu grâce à l'utilisation de plages.
	La plage appliquée à la tranche de destination signifie que les éléments copiés commenceront à la position 1.
	La plage appliquée à la tranche source signifie que la copie commencera par l'élément en position 2 et qu'un seul élément sera copié.
	**/
	product7s := [4]string{"kayak", "lifejacket", "paddle", "Hat"}
	allProduct7s := product7s[1:]
	someProduct7s := []string{"Boots", "Canoe"}
	copy(someProduct7s[1:], allProduct7s[2:3])
	fmt.Println("someProduct7s :", someProduct7s)
	fmt.Println("allProduct7s :", allProduct7s)

	/**
	Copie de tranches de tailles différentes
	Si la tranche de destination est plus grande que la tranche source, la copie se poursuivra jusqu'à ce que le dernier élément
	de la source ait été copié.

	La tranche source ne contient que deux éléments et aucune plage n'est utilisée. Par conséquent, la fonction `copy` commence à copier
	les éléments de la tranche replacementProduct8s vers la tranche product8s et s'arrête lorsque la fin de la tranche replacementProduct8s
	est atteinte. Les éléments restants dans la tranche product8s ne sont pas affectés par l'opération de copie.
	**/
	product8s := []string{"kayak", "lifejacket", "paddle", "Hat"}
	replacementProduct8s := []string{"Canoe", "Boots"}
	copy(product8s, replacementProduct8s)
	fmt.Println("product8s :", product8s)
	fmt.Println("replacementProduct8s :", replacementProduct8s)

	/**
	Si la tranche de destination est plus petite que la tranche source, la copie se poursuit jusqu'à ce que tous les éléments de la tranche
	de destination aient été remplacés.
	Dans l'exemple ci-dessous, la plage utilisée pour la destination crée une tranche de longueur 1, ce qui signifie qu'un seul élément
	sera copié du tableau source.
	**/
	product9s := []string{"kayak", "lifejacket", "paddle", "Hat"}
	replacementProduct9s := []string{"Canoe", "Boots"}
	copy(product9s[0:1], replacementProduct9s)
	fmt.Println("product9s :", product9s)
	fmt.Println("replacementProduct9s :", replacementProduct9s)
}
