package main

import (
	"fmt"
	_ "packages/data"
	. "packages/detail"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

/**
Les packages permettent de structurer un projet en regroupant les fonctionnalités apparentées.

Les dépendances aux packages personnalisés sont déclarées à l'aide de l'instruction `import`.
L'instruction d'importation spécifie le package sous forme de chemin, composé du nom du module (créé par la commande `go mod init modulename`)
et du nom du package, séparés par une barre oblique.

Les fonctionnalités exportées fournies par le package sont accessibles en utilisant le nom du package comme préfixe.
```
var product *store.Product = &store.Product{}
```

Lors de l'importation d'un paquet, la combinaison du nom du module et du nom du paquet garantit son identification unique.
Cependant, seul le nom du paquet est utilisé pour accéder à ses fonctionnalités, ce qui peut engendrer des conflits.
Prenons l'exemple avec le package 'fmt' créé. Le fichier `fmt/formats.go` exporte une fonction nommée ToCurrency qui reçoit une valeur
float64 et produit un montant en dollars formaté, à l'aide de la fonction `strconv.FormatFloat`.
Le package `fmt` défini porte le même nom qu'un des packages de la bibliothèque standard les plus utilisés.
Cela pose problème lorsque les deux packages sont utilisés simultanément. Une façon de gérer les conflits de noms de paquets consiste à
utiliser un alias, ce qui permet d'accéder à un paquet en utilisant un nom différent.
L'alias dans cet exemple résout le conflit de noms afin que les fonctionnalités définies par le package importé avec
le chemin packages/fmt soient accessibles en utilisant currencyFmt comme préfixe.
```
import (
	"fmt"
	currencyFmt "packages/fmt"
)
```

Il existe un alias spécial, appelé `dot import` (importation par points), qui permet d'utiliser les fonctionnalités d'un paquet
sans utiliser de préfixe. L'importation par point utilise un point (`.`) comme alias de package.
L'importation par point nous permet d'accéder à la fonction DisplayInfos (du package detail) sans utiliser de préfixe.
Lors de l'utilisation d'une importation pointée, nous devons nous assurer que les noms des fonctionnalités importées du package ne sont pas
définis dans le package importateur. Par exemple, cela signifie qu'on doit vérifier que le nom `DisplayInfos` n'est utilisé par aucune
fonctionnalité définie dans le package principal. C'est pourquoi les importations par point doivent être utilisées avec précaution.

Lors de l'importation d'un package imbriqué (`packages/store/cart`), le chemin du package commence par le nom du module et liste
la séquence des packages. Les fonctionnalités définies par le package imbriqué sont accessibles via le nom du package,
comme pour n'importe quel autre package.

Go empêche l'importation de packages sans leur utilisation, ce qui peut poser problème si vnous comptons sur l'effet d'une fonction
d'initialisation mais n'avez pas besoin d'utiliser les fonctionnalités exportées par le package.
Si nous avons besoin de l'effet de la fonction d'initialisation du package `packages/data` , mais que nous n'avons pas besoin
d'utiliser sa fonction `GetData` exportée par le package `packages/data`, nous pouvons importer le package en utilisant l'identifiant
vide comme alias du nom du package : `_ "packages/data"`.
L'identifiant vide (le caractère de soulignement) permet d'importer le package sans qu'il soit nécessaire d'utiliser ses fonctionnalités
exportées.
**/

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	fmt.Println("product Name :", product.Name)
	fmt.Println("product Category :", product.Category)
	fmt.Println("product Price :", product.Price())

	product1 := store.NewProduct("kayak", "Watersports", 0)
	product1.SetPrice(279.00)
	fmt.Println("product1 Name :", product1.Name)
	fmt.Println("product1 Category :", product1.Category)
	fmt.Println("product1 Price :", product1.Price())
	fmt.Println("product1 Price formated :", currencyFmt.ToCurrency(product1.Price()))
	fmt.Println("product1 detail :", DisplayInfos(product1.Name, product1.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{product, *product1},
	}
	fmt.Println("Cart name :", cart.CustomerName)
	fmt.Println("Cart total :", currencyFmt.ToCurrency(cart.GetTotal()))

	product2 := store.NewProduct("kayak", "Watersports", 275)
	fmt.Println("product2 Name :", product2.Name)
	fmt.Println("product2 Category :", product2.Category)
	fmt.Println("product2 Price :", product2.PriceWithCategory())
}
