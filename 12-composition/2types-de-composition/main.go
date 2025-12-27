package main

import (
	"composition/store"
	"fmt"
)

/**
Go prend en charge la composition, plutôt que l'héritage, qui se fait en combinant des types de structures.

Go permet d'accéder aux champs du type imbriqué de deux manières.
- La première est l'approche classique consistant à parcourir la hiérarchie des types pour atteindre la valeur recherchée.
Le champ `*Product` est imbriqué, ce qui signifie que son nom est son type.
`b.Product.Name`
- Le type `Boat` ne définit pas de champ `Name`, mais il peut être traité comme s'il en possédait un grâce à la fonctionnalité d'accès direct.
C'est ce qu'on appelle la `promotion de champ` ; Go aplatit donc les types de sorte que le type `Boat` se comporte comme s'il définissait
les champs fournis par le type `Product` imbriqué.

Les méthodes sont également promues afin que les méthodes définies pour le type imbriqué puissent être appelées depuis le type englobant.
Si le type de champ est une valeur, comme `Product`, toutes les méthodes définies avec des récepteurs `Product` ou `*Product` seront appelées.
Si le type de champ est un pointeur, comme `*Product`, seules les méthodes avec des récepteurs `*Product` seront appelées.
Dans notre exemple, il n'existe pas de méthode `Price` définie pour le type `*Boat`, mais Go promeut la méthode définie avec
un récepteur `*Product`.

Le type `RentalBoat` est composé à partir du type `*Boat`, lui-même composé à partir du type `*Product`, formant ainsi une chaîne.
Go effectue une promotion afin que les champs définis par les trois types de la chaîne soient directement accessibles.
Dans notre exemple, Go promeut les champs des types imbriqués Boat et Product afin qu'ils soient accessibles via le type de niveau supérieur
RentalBoat, ce qui permet de lire le champ Name. Les méthodes sont également promues au niveau supérieur.

Go ne peut effectuer la promotion que s'il n'existe aucun champ ou méthode portant le même nom sur le type englobant,
ce qui peut conduire à des résultats inattendus.
Dans notre exemple avec la structure `SpecialDeal`, les deux premiers résultats sont conformes à nos attentes :
les champs `Name` et `Price` du type `Product` ne sont pas mis en avant car le type `SpecialDeal` comporte des champs portant les mêmes noms.
Un problème similaire se pose lorsque deux champs imbriqués utilisent le même nom de champ ou de méthode.
```
type OfferBundle struct {
 	*store.SpecialDeal
 	*store.Product
}
```
**/

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println("Boat conventional access :", b.Product.Name, "- direct access :", b.Name)
	}

	for _, b := range boats {
		fmt.Println("Boat :", b.Name, "Price :", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
		store.NewRentalBoat("Yacht", 50000, 5, true, true),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat :", r.Name, "Rental Price :", r.Price(0.2))
	}

	luxuryRentalBoats := []*store.LuxuryRentalBoat{
		store.NewLuxuryRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewLuxuryRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewLuxuryRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, l := range luxuryRentalBoats {
		fmt.Println("Luxury rental Boat :", l.Name, "Luxury rental Price :", l.Price(0.2), "Captain :", l.Captain)
	}

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	dealName, dealPrice1, dealPrice2 := deal.GetDetails()
	fmt.Println("Deal Name :", dealName)
	fmt.Println("Deal Price field :", dealPrice1)
	fmt.Println("Deal Price method :", dealPrice2)
}
