package main

import "time"

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
}
