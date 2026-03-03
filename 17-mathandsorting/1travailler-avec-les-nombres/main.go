package main

import "math"

/**
Le langage Go prend en charge un ensemble d'opérateurs arithmétiques applicables aux valeurs numériques, permettant d'effectuer des opérations de
base telles que l'addition et la multiplication. Pour des opérations plus avancées, la bibliothèque standard de Go inclut le package `math`, qui
offre un large éventail de fonctions.

Abs(val) : Cette fonction renvoie la valeur absolue d'une valeur float64, c'est-à-dire la distance à zéro sans tenir compte du sens.

Ceil(val) : Cette fonction renvoie le plus petit entier supérieur ou égal à la valeur float64 spécifiée.
Le résultat est également une valeur float64, même s'il représente un entier.

Copysign(x, y) : Cette fonction renvoie une valeur float64 qui correspond à la valeur absolue de x avec le signe de y.

Floor(val) : Cette fonction renvoie le plus grand entier inférieur ou égal à la valeur float64 spécifiée.
Le résultat est également une valeur float64, même s'il représente un entier.

Max(x, y) : Cette fonction renvoie la plus grande des valeurs float64 spécifiées.

Min(x, y) : Cette fonction renvoie la plus petite des valeurs float64 spécifiées.

Mod(x, y) : Cette fonction renvoie le reste de la division de x par y.

Pow(x, y) : Cette fonction renvoie x élevé à la puissance y.

Round(val) : Cette fonction arrondit la valeur spécifiée à l'entier le plus proche, en arrondissant les demi-valeurs à l'entier supérieur.
Le résultat est une valeur float64, même si elle représente un entier.

RoundToEven(val) : Cette fonction arrondit la valeur spécifiée à l'entier le plus proche, en arrondissant les demi-valeurs à l'entier pair
le plus proche. Le résultat est une valeur float64, même si elle représente un entier.

Ces fonctions opèrent toutes sur des valeurs float64 et produisent des résultats float64, ce qui signifie que nous devons explicitement
effectuer des conversions vers et depuis d'autres types.

Le package math fournit également un ensemble de constantes pour les limites des types de données numériques :
MaxInt8
MinInt8 : Ces constantes représentent les valeurs maximale et minimale pouvant être stockées dans un entier 8 bits (int8).

MaxInt16
MinInt16 : Ces constantes représentent les valeurs maximale et minimale pouvant être stockées dans un entier 16 bits (int16).

MaxInt32
MinInt32 : Ces constantes représentent les valeurs maximale et minimale pouvant être stockées dans un entier 32 bits (int32).

MaxInt64
MinInt64 : Ces constantes représentent les valeurs maximale et minimale pouvant être stockées dans un entier 64 bits (int64).

MaxUint8 : Cette constante représente la valeur maximale pouvant être représentée dans un entier non signé 8 bits (uint8).
La valeur minimale est zéro.

MaxUint16 : Cette constante représente la valeur maximale pouvant être représentée dans un entier non signé 16 bits (uint16).
La valeur minimale est zéro.

MaxUint32 : Cette constante représente la valeur maximale pouvant être représentée dans un entier non signé 32 bits (uint32).
La valeur minimale est zéro. MaxUint64 : Cette constante représente la plus grande valeur pouvant être représentée par un entier non signé 64 bits (uint64). La plus petite valeur est zéro.

MaxFloat32
MaxFloat64 : Ces constantes représentent les plus grandes valeurs pouvant être représentées par des nombres à virgule flottante
32 bits (float32) et 64 bits (float64).

SmallestNonzeroFloat32
SmallestNonzeroFloat32 : Ces constantes représentent les plus petites valeurs non nulles pouvant être représentées par des nombres à virgule flottante
32 bits (float32) et 64 bits (float64).
**/

func main() {
	val1 := 279.00
	val2 := 48.95

	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
}
