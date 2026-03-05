package main

import (
	"fmt"
	"time"
)

/**
Le package `time` offre des fonctionnalités permettant de mesurer les durées et d'exprimer les dates et les heures.
Le package `time` fournit le type `Time`, utilisé pour représenter un instant précis.

Fonctions du package `time` pour la création de valeurs de type `Time`

Now() : Cette fonction crée une valeur de type `Time` représentant l'instant présent.

Date(y, m, d, h, min, sec, nsec, loc) : Cette fonction crée une valeur de type `Time` représentant un instant précis, exprimé par l'année,
le mois, le jour, l'heure, la minute, la seconde, la nanoseconde et le lieu (loc).
Le dernier argument de la fonction `Date` est un emplacement, qui spécifie le lieu dont le fuseau horaire sera utilisé pour la valeur de l'heure.

Unix(sec, nsec) : Cette fonction crée une valeur de type `Time` correspondant au nombre de secondes et de nanosecondes écoulées depuis le
1er janvier 1970 UTC, communément appelé temps Unix.


Un objet de type `Time` contient plusieurs informations. Pour récupérer ces informations, on utilise des méthodes fournies par ce type :

Date() : Cette méthode renvoie l'année, le mois et le jour. L'année et le jour sont exprimés sous forme d'entiers (int) et le mois sous forme
d'objet `Month`.

Clock() : Cette méthode renvoie l'heure, les minutes et les secondes.

Year() : Cette méthode renvoie l'année, exprimée sous forme d'entier (int).

YearDay() : Cette méthode renvoie le jour de l'année, exprimé sous forme d'entier (int) compris entre 1 et 366 (pour tenir compte des années bissextiles).

Month() : Cette méthode renvoie le mois, exprimé sous forme d'objet `Month`.

Day() : Cette méthode renvoie le jour du mois, exprimé sous forme d'entier (int).

Weekday() : Cette méthode renvoie le jour de la semaine, exprimé sous forme d'objet Weekday.

Hour() : Cette méthode renvoie l'heure du jour, sous forme d'entier compris entre 0 et 23.

Minute() : Cette méthode renvoie le nombre de minutes écoulées avant l'heure du jour, sous forme d'entier compris entre 0 et 59.

Second() : Cette méthode renvoie le nombre de secondes écoulées avant la minute de l'heure, sous forme d'entier compris entre 0 et 59.

Nanosecond() : Cette méthode renvoie le nombre de nanosecondes écoulées avant la seconde de la minute, sous forme d'entier compris
entre 0 et 999 999 999.


Deux types sont définis pour aider à décrire certaines composantes d'une valeur de type `Time` :

Month: Ce type représente un mois. Le package `time` définit des valeurs constantes pour les noms de mois en anglais : janvier, février, etc.
Le type `Month` définit une méthode `String` qui utilise ces noms lors de la mise en forme des chaînes de caractères.

Weekday : Ce type représente un jour de la semaine. Le package `time` définit des valeurs constantes pour les noms de jours de la semaine en anglais :
Sunday, Monday, etc. Le type `Weekday` définit une méthode `String` qui utilise ces noms lors de la mise en forme des chaînes de caractères.


Formatage des dates sous forme de chaînes de caractères

La méthode `Format` permet de créer des chaînes de caractères formatées à partir de valeurs de type `Time`. Le format de la chaîne est spécifié par
une chaîne `layout`, qui indique les composants de l'heure requis, ainsi que leur ordre et leur précision.
Format(layout) : Cette méthode renvoie une chaîne formatée, créée à l'aide de la mise en page spécifiée.

La chaîne `layout` utilise une heure de référence, qui est 15:04:05 (soit cinq secondes après quatre minutes après 15 heures)
le lundi 2 janvier 2006, dans le fuseau horaire MST, qui est 7 heures en retard sur l'heure moyenne de Greenwich (GMT).
Dans le layout de Format, on peut écrire à la fois des éléments représentant la date/heure et du texte normal :
- t.Format("2006-01-02")
- t.Format("02/01/2006 à 15:04")

Le package `time` définit un ensemble de constantes pour les formats de date et d'heure courants :
(Ces constantes peuvent être utilisées à la place d'une mise en page personnalisée.)

ANSIC        : Mon Jan _2 15:04:05 2006
UnixDateMon  : Jan _2 15:04:05 MST 2006
RubyDateMon  : Jan 02 15:04:05 -0700 2006
RFC82202     : Jan 06 15:04 MST
RFC822Z02    : Jan 06 15:04 -0700
RFC850       : Monday, 02-Jan-06 15:04:05 MST
RFC1123      : Mon, 02 Jan 2006 15:04:05 MST
RFC1123Z     : Mon, 02 Jan 2006 15:04:05 -0700
RFC3339      : 2006-01-02T15:04:05Z07:00
RFC3339Nano  : 2006-01-02T15:04:05.999999999Z07:00
Kitchen      : 3:04PM
Stamp        : Jan _2 15:04:05
StampMilli   : Jan _2 15:04:05.000
StampMicro   : Jan _2 15:04:05.000000
StampNano    : Jan _2 15:04:05.000000000
**/

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
}

func PrintFormatedTime(label string, t *time.Time) {
	layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(layout))
}

func PrintTimeWithPredefinedLayout(label string, t *time.Time) {
	fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)

	Printfln("\nFormated time with custom layout \n")

	PrintFormatedTime("Current", &current)
	PrintFormatedTime("Specific", &specific)
	PrintFormatedTime("UNIX", &unix)

	Printfln("\nFormated time with predefined layout \n")

	PrintTimeWithPredefinedLayout("Current", &current)
	PrintTimeWithPredefinedLayout("Specific", &specific)
	PrintTimeWithPredefinedLayout("UNIX", &unix)
}
