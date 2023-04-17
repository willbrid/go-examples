package main

/** Une façon de gérer les conflits de noms de packages consiste à utiliser un alias,
ce qui permet d'accéder à un package en utilisant un nom différent.
 Il existe un alias spécial, connu sous le nom d'importation de points, qui permet d'utiliser les fonctionnalités d'un package sans utiliser de préfixe

Si nous avons besoin de l'effet de la fonction d'initialisation du package data, mais que
nous n'avons pas besoin d'utiliser la fonction GetData que le package exporte,
nous pouvons importer le package en utilisant l'identifiant vide (_) comme alias pour le nom du package.
 **/
import (
	"fmt"
	_ "packages/data"
	currencyFmt "packages/fmt" // Alias : currencyFmt
	. "packages/greet"         // Importation de points
	"packages/store"           // On précise le module puis le nom du package
	"packages/store/cart"      // Importation du package imbriqué

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello, Packages and Modules")

	// Création d'un custom package
	fmt.Println("Création d'un custom package")

	var product1 store.Product = store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	fmt.Println("Name : ", product1.Name)
	fmt.Println("Category:", product1.Category)

	var product2 *store.Product = store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name : ", product2.Name)
	fmt.Println("Category : ", product2.Category)
	fmt.Println("Price : ", product2.Price())
	fmt.Println("Price standard : ", product2.PriceStandardTax())

	// Dans le même package, les noms de classe doivent être unique.
	var product3 *store.Product = store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Price:", currencyFmt.ToCurrency(product3.Price()))

	// Lors de l'utilisation d'une importation de points, nous devons vous assurer que les noms des entités importées
	// à partir du package ne sont pas définis dans le package d'importation.
	var greeting string = GetHello("Willbrid")
	fmt.Println("Message : ", greeting)

	// Utilisation du package imbriqué
	var cart cart.Cart = cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product2, *product3},
	}
	fmt.Println("Name : ", cart.CustomerName)
	fmt.Println("Total : ", currencyFmt.ToCurrency(cart.GetTotal()))
	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))
}
