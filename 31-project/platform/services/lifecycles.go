package services

/**
Les cycles de vie des services :
- Transient : pour ce cycle de vie, la fonction factory (usine) est appelée pour chaque demande de service.
- Singleton : pour ce cycle de vie, la fonction factory (usine) est invoquée une fois et chaque requête reçoit la même instance de classe.
- Scoped : pour ce cycle de vie, la fonction factory (usine) est appelée une fois pour la première demande dans une portée, et chaque demande
           dans cette portée reçoit la même instance de classe.
**/

type lifecycle int

const (
	Transient lifecycle = iota
	Singleton
	Scoped
)
