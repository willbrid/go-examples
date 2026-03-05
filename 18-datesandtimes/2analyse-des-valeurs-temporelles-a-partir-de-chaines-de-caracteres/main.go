package main

import (
	"fmt"
	"time"
)

/**
Le package `time` permet de créer des valeurs `Time` à partir de chaînes de caractères.
Fonctions du package time pour l'analyse des chaînes de caractères en valeurs de type `Time`

Parse(layout, str) : Cette fonction analyse une chaîne de caractères en utilisant le `layout` spécifié pour créer une valeur de type `Time`.
Une erreur est renvoyée en cas de problème d'analyse.

ParseInLocation(layout, str, location) : Cette fonction analyse une chaîne de caractères en utilisant le `layout` spécifié et
la localisation si la chaîne ne contient pas de fuseau horaire. Une erreur est renvoyée en cas de problème d'analyse.


Ces fonctions utilisent un temps de référence, servant à spécifier le format de la chaîne à analyser. Ce temps de référence est
15 h 04 min 05 s (soit cinq secondes après 15 h 04) le lundi 2 janvier 2006, heure normale des Rocheuses (MST), qui a sept heures de retard sur
le temps universel coordonné (UTC).


Spécification d'un localisation parsée
La fonction `Parse` suppose que les dates et heures exprimées sans fuseau horaire sont définies en temps universel coordonné (UTC).
La méthode `ParseInLocation` permet de spécifier un emplacement qui sera utilisé lorsqu'aucun fuseau horaire n'est indiqué.
La fonction ParseInLocation accepte en 3ème argument `time.Location` qui spécifie un lieu dont le fuseau horaire sera utilisé
lorsqu'il n'est pas inclus dans la chaîne analysée. Les valeurs de type `time.Location` peuvent être créées à l'aide des fonctions :

- LoadLocation(name) : Cette fonction renvoie une valeur *Location correspondant au nom spécifié, ainsi qu'une erreur indiquant tout problème rencontré.

- LoadLocationFromTZData(name, data) : Cette fonction renvoie une valeur *Location à partir d'une slice d'octets contenant une base
de données de fuseaux horaires formatée.

- FixedZone(name, offset) : Cette fonction renvoie une valeur *Location utilisant systématiquement le nom et le décalage par rapport à UTC spécifiés.

Lorsqu'un lieu est transmis à la fonction `LoadLocation`, la propriété Location renvoyée contient les détails des fuseaux horaires utilisés à
cet endroit. Les noms de lieux sont définis dans la base de données de fuseaux horaires de l'IANA (https://www.iana.org/time-zones) et sont
répertoriés sur la page Wikipédia suivante : https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.

La base de données de fuseaux horaires utilisée pour créer les valeurs de localisation est installée avec les outils Go, ce qui signifie qu'elle
peut ne pas être disponible lors du déploiement d'une application compilée. Le package `time/tzdata` contient une version intégrée de la
base de données, chargée par une fonction d'initialisation du package. Pour garantir la disponibilité permanente des données de fuseaux horaires,
l'on doit déclarer une dépendance au package comme ceci :

import (
  ...
  _ "time/tzdata"
  ...
)

Utilisation de la localisation locale
Si le nom de lieu utilisé pour créer une localisation est « Local », le fuseau horaire de la machine exécutant l’application est utilisé.
time.LoadLocation("Local")


Spécifier directement les fuseaux horaires
L'utilisation des noms de lieux est la méthode la plus fiable pour garantir l'exactitude des dates, car l'heure d'été est automatiquement appliquée.
La fonction `FixedZone` permet de créer une valeur de type `*Location` avec un fuseau horaire fixe.
**/

func PrintTime(label string, t *time.Time) {
	fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}

	for _, d := range dates {
		timeVal, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed :", &timeVal)
		} else {
			Printfln("Error : %s", err.Error())
		}
	}

	/**
	Spécification d'un localisation parsée
	**/

	layout1 := "02 Jan 06 15:04"
	date1 := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")

	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout1, date1) // Timezone : UTC (+0000 à la fin)
		londonTime, _ := time.ParseInLocation(layout1, date1, london)
		newyorkTime, _ := time.ParseInLocation(layout1, date1, newyork)
		localTime, _ := time.ParseInLocation(layout1, date1, local)

		PrintTime("No location :", &nolocation)
		PrintTime("London :", &londonTime)
		PrintTime("Newyork :", &newyorkTime)
		PrintTime("Local :", &localTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}

	/**
	Spécifier directement les fuseaux horaires
	Les arguments de la fonction FixedZone sont un nom et le nombre de secondes de décalage par rapport à UTC.
	Cet exemple crée trois fuseaux horaires fixes :
	- l’un est en avance d’une heure sur UTC,
	- l’autre en retard de quatre heures
	- le dernier n’a pas de décalage.
	**/

	layout2 := "02 Jan 06 15:04"
	date2 := "09 Jun 95 19:30"
	london1 := time.FixedZone("BST", 1*60*60)
	newyork1 := time.FixedZone("EDT", -4*60*60)
	local1 := time.FixedZone("Local", 0)

	nolocation1, _ := time.Parse(layout2, date2) // Timezone : UTC (+0000 à la fin)
	londonTime1, _ := time.ParseInLocation(layout2, date2, london1)
	newyorkTime1, _ := time.ParseInLocation(layout2, date2, newyork1)
	localTime1, _ := time.ParseInLocation(layout2, date2, local1)

	PrintTime("No location1 :", &nolocation1)
	PrintTime("London1 :", &londonTime1)
	PrintTime("Newyork1 :", &newyorkTime1)
	PrintTime("Local1 :", &localTime1)
}
