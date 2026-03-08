package main

import (
	"fmt"
	"time"
)

/**
Le type `Duration` est un alias du type int64 et sert à représenter un nombre précis de millisecondes. Les valeurs `Duration` personnalisées sont
composées de valeurs `Duration` constantes définies dans le package time.

`Hour` : Cette constante représente 1 heure.
`Minute` : Cette constante représente 1 minute.
`Second` : Cette constante représente 1 seconde.
`Millisecond` : Cette constante représente 1 milliseconde.
`Microsecond` : Cette constante représente 1 microseconde.
`Nanosecond` : Cette constante représente 1 nanoseconde.

Une fois une valeur de type `Duration créée, elle peut être inspectée à l'aide des méthodes (méthode du type `Duration`) :
(Ces méthodes renvoient la durée totale exprimée dans une unité spécifique, comme les heures ou les minutes.
Cela diffère des méthodes aux noms similaires définies par le type `Time`, qui ne renvoient qu'une partie de la date/heure.)

Hours() : Cette méthode renvoie un nombre flottant (float64) représentant la durée en heures.

Minutes() : Cette méthode renvoie un nombre flottant (float64) représentant la durée en minutes.

Seconds() : Cette méthode renvoie un nombre flottant (float64) représentant la durée en secondes.

Milliseconds() : Cette méthode renvoie un entier (int64) représentant la durée en millisecondes.

Microseconds() : Cette méthode renvoie un entier (int64) représentant la durée en microsecondes.

Nanoseconds() : Cette méthode renvoie un entier (int64) représentant la durée en nanosecondes.

Round(duration) : Cette méthode renvoie une durée arrondie au multiple le plus proche de la durée spécifiée.

Truncate(duration) : Cette méthode renvoie une durée arrondie à l'entier inférieur de la durée spécifiée.


Création de durées relatives à une heure
Le package `time` définit deux fonctions permettant de créer des valeurs de durée représentant l'intervalle de temps entre une
heure spécifique et l'heure actuelle.

`Since(time)` : Cette fonction renvoie une durée exprimant le temps écoulé depuis la valeur de `time` spécifiée.

`Until(time)` : Cette fonction renvoie une durée exprimant le temps écoulé jusqu'à la valeur de `time` spécifiée.


Création de durées à partir de chaînes de caractères
La fonction `time.ParseDuration` analyse les chaînes de caractères pour créer des valeurs `Duration`.

ParseDuration(str) : Cette fonction renvoie une durée et une erreur, indiquant s'il y a eu des problèmes lors de l'analyse de la chaîne spécifiée.

Le format des chaînes de caractères prises en charge par la fonction `ParseDuration` est une séquence de valeurs numériques suivie
des indicateurs d'unité ci-dessous :

h        : Cette unité désigne les heures.
m        : Cette unité désigne les minutes.
s        : Cette unité désigne les secondes.
ms       : Cette unité désigne les millisecondes.
us ou µs : Ces unités désignent les microsecondes.
ns       : Cette unité désigne les nanosecondes.

Aucun espace n'est autorisé entre les valeurs, qui peuvent être spécifiées sous forme d'entiers ou de nombres à virgule flottante.
**/

func main() {
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())

	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())

	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())

	// L'exemple utilise les méthodes `Until` et `Since` pour calculer le nombre d'années restantes jusqu'en 2051 et le nombre d'années écoulées depuis 1965.
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))

	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}
