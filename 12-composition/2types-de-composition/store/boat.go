package store

/**
Le type de structure `Boat` définit un champ `*Product` intégré.

Une structure peut combiner des champs classiques et des champs imbriqués, mais ces derniers jouent un rôle important dans la composition.
La fonction `NewBoat` est un constructeur qui utilise ses paramètres pour créer un objet `Boat`, avec sa valeur `Product` imbriquée.
**/

type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}
