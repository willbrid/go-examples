package store

/**
Les types peuvent définir plusieurs champs de structure, et Go promouvra ces champs pour chacun d'eux.
Le type `LuxuryRentalBoat` possède des champs `*Boat` et `*Crew`, et Go promeut les champs et les méthodes des deux types imbriqués.
**/

type LuxuryRentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

func NewLuxuryRentalBoat(name string, price float64, capacity int, motorized, crew bool, captain, firstOffice string) *LuxuryRentalBoat {
	return &LuxuryRentalBoat{
		NewBoat(name, price, capacity, motorized),
		crew,
		NewCrew(captain, firstOffice),
	}
}
