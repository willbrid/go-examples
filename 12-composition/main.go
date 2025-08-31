package main

import (
	"composition/store"
	"fmt"
)

func main() {
	fmt.Println("Hello, Composition")

	// Comprendre la composition des types
	fmt.Println("Comprendre la composition des types")
	var kayak *store.Product = store.NewProduct("Kayak", "Watersports", 275)
	var lifejacket *store.Product = &store.Product{Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name : ", p.Name, " - Category : ", p.Category, " - Price : ", p.Price(0.2))
	}

	var boats []*store.Boat = []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	/**
	Pour atteindre le champ Name, nous pouvons naviguer dans le type imbriqué Product.

	Go permet également d'utiliser directement les types de champs imbriqués. Ainsi le type Boat ne définit pas de champ Name,
	mais il peut être traité comme s'il le faisait grâce à la fonction d'accès direct.

	Les méthodes sont également promues afin que les méthodes définies pour
	le type imbriqué (Product) puissent être appelées à partir du type englobant (Boat).

	On a deux types de promotion : promotion de champ et promotion de méthode.
	**/
	for _, b := range boats {
		fmt.Println("Conventional : ", b.Product.Name, " - Direct : ", b.Name, " - Price : ", b.Price(0.2))
	}

	/** Go permet d'initialiser directement un type composée. Cependant il est plus flexible d'utiliser les fonctions constructeur et
	d'invoquer un constructeur à partir d'un autre, comme par exemple le constructeur NewBoat qui appelle le constructeur NewProduct
	**/
	var boat store.Boat = store.Boat{Product: &store.Product{Name: "Kayak", Category: "Watersports"}, Capacity: 1, Motorized: false}
	fmt.Println("Conventional : ", boat.Product.Name, " - Direct : ", boat.Name, " - Price : ", boat.Price(0.2))

	var rentals []*store.RentalBoat = []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat : ", r.Name, " - Rental Price : ", r.Price(0.2), " - Captain : ", r.Captain)
	}

	/**
	Les champs Name et price du type de Product ne sont pas promus car le type SpecialDeal a des champs avec les mêmes noms : Name et price.
	Afin d'utiliser le champ price de SpecialDeal, on définit une nouvelle méthode appelée Price de SpecialDeal qui utilise le champ price de SpecialDeal.
	Ce qui va empêcher la promotion de la méthode Price du type Product depuis le type SpecialDeal.
	**/
	var product *store.Product = store.NewProduct("Kayak", "Watersports", 279)
	var deal *store.SpecialDeal = store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price1, price2 := deal.GetDetails()
	fmt.Println("Name : ", Name)
	fmt.Println("Price field : ", price1)
	fmt.Println("Price method : ", price2)
	/**
	Un autre problème connexe survient lorsque deux champs imbriqués utilisent le même nom de champ ou de méthode.
	Le type composé est distinct et ne peut pas être utilisé lorsque les types à partir desquels il est composé sont requis.

	Changer le type map pour qu'elle utilise l'interface nous permet de stocker les valeurs de Product et de Boat. Le type Product
	se conforme directement à l'interface ItemForSale car il existe une méthode Price qui correspond à la signature spécifiée par
	l'interface et qui a un récepteur *Product.
	Il n'y a pas de méthode Price qui utilise un récepteur *Boat, mais Go prend en compte la méthode Price promue à partir du champ
	intégré du type Boat, qu'elle utilise pour satisfaire aux exigences de l'interface.
	**/
	var products map[string]store.ItemForSale = map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for key, p := range products {
		fmt.Println("Key : ", key, " - Price : ", p.Price(0.2))
	}
	/**
	Une assertion de type est effectuée par l'instruction "case" du "switch" lorsqu'un seul type est spécifié, même si cela peut conduire à
	une duplication lorsque chaque type est traité.
	**/
	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name : ", item.Name, " - Category : ", item.Category, " - Price : ", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name : ", item.Name, " - Category : ", item.Category, " - Price : ", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	/**
	Composition d'interface : une interface peut en enfermer une autre, avec pour effet que les types doivent implémenter
	toutes les méthodes définies par les interfaces englobantes et enfermées. Les interfaces sont plus simples que les structures,
	et il n'y a pas de champs ou de méthode à promouvoir. Le résultat de la composition des interfaces est une union de la méthode
	définie par les types englobant et clos.
	**/
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			// Plus nécessaire de passer par ItemForSale car il est déjà une interface composite de Describable
			fmt.Println("Name : ", item.GetName(), " - Category : ", item.GetCategory(), " - Price : ", item.(store.ItemForSale).Price(0.2))
			fmt.Println("Name : ", item.GetName(), " - Category : ", item.GetCategory(), " - Price : ", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

}
