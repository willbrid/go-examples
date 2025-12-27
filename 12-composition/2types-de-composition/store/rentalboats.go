package store

/**
La fonctionnalité de composition peut être utilisée pour créer des chaînes complexes de types imbriqués, dont les champs et les méthodes
sont promus au type englobant de niveau supérieur.
**/

type RentalBoat struct {
	*Boat
	IncludeCrew bool
}

func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool) *RentalBoat {
	return &RentalBoat{
		NewBoat(name, price, capacity, motorized),
		crewed,
	}
}
