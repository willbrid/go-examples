package main

import "time"

func PrintTime(label string, t *time.Time) {
	/**
	Cette méthode t.Day() (time.Time) renvoie le jour du mois, exprimé sous la forme d'un entier.
	Cette méthode t.Month() (time.Time) renvoie le composant mois, exprimé à l'aide du type Mois.
	Cette méthode t.Year() (time.Time) renvoie le composant de l'année, exprimé sous la forme d'un int.
	Cette méthode t.Hour() renvoie l'heure du jour, exprimée sous la forme d'un entier compris entre 0 et 23.
	Cette méthode t.Minute() renvoie le nombre de minutes écoulées dans l'heure du jour, exprimée sous la forme d'un entier compris entre 0 et 59.
	Cette méthode t.Second() renvoie le nombre de secondes écoulées dans la minute de l'heure, exprimée sous la forme d'un entier compris entre 0 et 59.
	Cette méthode t.Nanosecond() renvoie le nombre de nanosecondes écoulées dans la seconde de la minute, exprimée sous la forme d'un entier compris entre 0 et 999 999 999.
	**/
	Printfln("%s : Day: %v - Month : %v - Year : %v, Hour : %v - Minute : %v - Seconde : %v - Nanoseconde : %v", label, t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

func PrintTime1(label string, t *time.Time) {
	var layout string = "Day: 02 Month: Jan Year: 2006"
	Printfln("1- %v : %v", label, t.Format(layout))
}

func PrintTime2(label string, t *time.Time) {
	/**
	Le package time définit un ensemble de constantes pour les formats d'heure et de date courants.
	Pour ce cas nous avons utilisé le format time.RFC822Z
	**/
	Printfln("2- %v : %v", label, t.Format(time.RFC822Z))
}

func main() {
	Printfln("Hello, Dates and Times")

	// Cette fonction time.Now() crée un Time représentant le moment actuel dans le temps.
	current := time.Now()
	/**
	  Cette fonction time.Date() crée un Time représentant un moment spécifié dans le temps, qui est exprimé par
	  les arguments year, month, day, hour, minute, second, nanosecond
	  et Location (Local représente le fuseau horaire local du système. Sur les systèmes Unix, time.Local consulte la variable d'environnement
	  TZ pour trouver le fuseau horaire à utiliser.).
	**/
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	/**
	  Cette fonction time.Unix() crée une valeur de temps à partir du nombre de secondes et de nanosecondes depuis le 1er janvier 1970, UTC,
	  communément appelé temps Unix.
	**/
	unix := time.Unix(1433228090, 0)

	PrintTime("Current : ", &current)
	PrintTime("Specific : ", &specific)
	PrintTime("Unix : ", &unix)

	current1 := time.Now()
	specific1 := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix1 := time.Unix(1433228090, 0)
	PrintTime1("Current1 : ", &current1)
	PrintTime1("Specific1 : ", &specific1)
	PrintTime1("Unix1 : ", &unix1)

	current2 := time.Now()
	specific2 := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix2 := time.Unix(1433228090, 0)
	PrintTime2("Current2 : ", &current2)
	PrintTime2("Specific2 : ", &specific2)
	PrintTime2("Unix2 : ", &unix2)

	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		/**
		Cette fonction time.Parse() analyse une chaîne à l'aide de la disposition spécifiée pour créer une valeur Time.
		Une erreur est renvoyée pour indiquer des problèmes d'analyse de la chaîne.
		**/
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime2("Parsed : ", &time)
		} else {
			Printfln("Error : %s", err.Error())
		}
	}

	date1s := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}
	for _, d := range date1s {
		// Cet exemple utilise la constante time.RFC822 pour analyser les chaînes de date.
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTime2("Parsed : ", &time)
		} else {
			Printfln("Error : %s", err.Error())
		}
	}

	layout1 := "02 Jan 06 15:04"
	date1 := "09 Jun 95 19:30"
	// Cette fonction time.LoadLocation() renvoie un *time.Location pour le nom spécifié et un erreur (error) qui indique tout problème.
	london, lonerr := time.LoadLocation("Europe/London")
	douala, douerr := time.LoadLocation("Africa/Douala")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	if lonerr == nil && douerr == nil && nycerr == nil {
		// Lorsque la méthode time.Parse est utilisée, le fuseau horaire est supposé être UTC, qui a un décalage null.
		noLocation, _ := time.Parse(layout1, date1)
		/**
		Cette fonction time.ParseInLocation() analyse une chaîne, en utilisant la disposition spécifiée et en utilisant l'emplacement
		si aucun fuseau horaire n'est inclus dans la chaîne. Une erreur est renvoyée pour indiquer des problèmes d'analyse de la chaîne.
		**/
		londonTime, _ := time.ParseInLocation(layout1, date1, london)
		doualaTime, _ := time.ParseInLocation(layout1, date1, douala)
		newyorkTime, _ := time.ParseInLocation(layout1, date1, newyork)
		// Si le nom de lieu utilisé pour créer un emplacement est Local, le paramètre de fuseau horaire de la machine exécutant l'application est utilisé.
		localTime, _ := time.ParseInLocation(layout1, date1, local)

		PrintTime2("No location : ", &noLocation)
		PrintTime2("London : ", &londonTime)
		PrintTime2("Douala : ", &doualaTime)
		PrintTime2("New York : ", &newyorkTime)
		PrintTime2("Local : ", &localTime)
	} else {
		Printfln(lonerr.Error(), douerr.Error(), nycerr.Error())
	}

	// Cette fonction time.FixedZone renvoie un pointeur *Location qui utilise toujours le nom spécifié et le décalage par rapport à UTC.
	london1 := time.FixedZone("BST", 1*60*60)
	newyork1 := time.FixedZone("EDT", -4*60*60)
	local1 := time.FixedZone("Local", 0)
	londonTime1, _ := time.ParseInLocation(layout1, date1, london1)
	newyorkTime1, _ := time.ParseInLocation(layout1, date1, newyork1)
	localTime1, _ := time.ParseInLocation(layout1, date1, local1)
	PrintTime2("London : ", &londonTime1)
	PrintTime2("New York : ", &newyorkTime1)
	PrintTime2("Local : ", &localTime1)

	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		// Cette méthode t.After renvoie true si l'heure à laquelle la méthode a été appelée se produit après l'heure fournie en argument.
		Printfln("After : %v", t.After(time.Now()))
		// Cette méthode t.Before renvoie true si l'heure à laquelle la méthode a été appelée se produit avant l'heure fournie en argument.
		Printfln("Before : %v", t.Before(time.Now()))
		// Cette méthode t.Round arrondit le temps à l'intervalle le plus proche représenté par une valeur de durée.
		Printfln("Round : %v", t.Round(time.Hour))
		// Cette méthode t.Truncate arrondit le temps à l'intervalle inférieur le plus proche représenté par une valeur de durée.
		Printfln("Truncate : %v", t.Truncate(time.Hour))
	} else {
		Printfln("Error : %v", err.Error())
	}

	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	// Cette méthode Equal renvoie true si l'heure à laquelle la méthode a été appelée est égale à l'heure fournie en argument.
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
	/**
	  Cette méthode t1.IsZero renvoie true si l'heure à laquelle la méthode a été appelée représente l'instant zéro,
	  qui est le 1er janvier, année 1, 00:00:00 UTC.
	**/
	Printfln("Is zero : %v", t1.IsZero())
	// Cette méthode t2.Location renvoie l'emplacement qui est associé à l'heure, permettant ainsi d'exprimer une heure dans un fuseau horaire différent.
	Printfln("Location : %v", t2.Location().String())

	var d time.Duration = time.Hour + (30 * time.Minute)
	// Cette méthode d.Hours renvoie un float64 qui représente la durée en heures.
	Printfln("Hours : %v", d.Hours())
	// Cette méthode d.Minutes renvoie un float64 qui représente la durée en minutes.
	Printfln("Mins : %v", d.Minutes())
	// Cette méthode d.Seconds renvoie un float64 qui représente la durée en secondes.
	Printfln("Seconds : %v", d.Seconds())
	// Cette méthode d.Milliseconds renvoie un int64 qui représente la durée en millisecondes.
	Printfln("Millseconds : %v", d.Milliseconds())
	// Cette méthode d.Round renvoie une durée, qui est arrondie au multiple le plus proche de la durée spécifiée.
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	// Cette méthode d.Truncate renvoie une durée, qui est arrondie au multiple inférieur le plus proche de la durée spécifiée.
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())

	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	// Cette fonction time.Until renvoie une durée exprimant le temps écoulé jusqu'à la valeur de temps spécifiée.
	Printfln("Futur : %v", toYears(time.Until(future)))
	// Cette fonction time.Since renvoie une durée exprimant le temps écoulé depuis la valeur de temps spécifiée.
	Printfln("Past : %v", toYears(time.Since(past)))

	/**
		h : Cette unité désigne les heures.
		m : Cette unité désigne les minutes.
		s : Cette unité désigne les secondes.
		ms : Cette unité désigne les millisecondes.
		us ou μs : Cette unité désigne les microsecondes.
		ns : Cette unité désigne les nanosecondes.
	**/
	d1, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d1.Hours())
		Printfln("Mins: %v", d1.Minutes())
		Printfln("Seconds: %v", d1.Seconds())
		Printfln("Millseconds: %v", d1.Milliseconds())
	} else {
		Printfln("Error : %v", err.Error())
	}
}
