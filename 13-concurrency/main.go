package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main function started")
	fmt.Println("Day : ", time.Now().Day(), "-", time.Now().Month(), "-", time.Now().Year())
	CalcStoreTotal(Products)
	// time.Sleep(time.Second * 5) // On retarde la goroutine principale pour permettre aux autres goroutines de terminer leur exécution
	fmt.Println("main function complete")
	/**
		Le bloc de construction clé pour l'exécution d'un programme Go est la goroutine, qui est un thread léger créé par le runtime Go.
		Tous les programmes Go utilisent au moins une goroutine car c'est ainsi que Go exécute le code dans la fonction main.
		Lorsque le code Go compilé est exécuté, le runtime crée une goroutine qui commence à exécuter les instructions dans le point d'entrée,
		qui est la fonction main du package main. Chaque instruction de la fonction main est exécutée dans l'ordre dans lequel elle est définie.
		La goroutine continue d'exécuter des instructions jusqu'à ce qu'elle atteigne la fin de la fonction main, moment auquel l'application se termine.
		La goroutine exécute chaque instruction de la fonction main de manière synchrone, ce qui signifie qu'elle attend que l'instruction se
		termine avant de passer à l'instruction suivante.

		Une goroutine est créée à l'aide du mot-clé go suivi de la fonction ou de la méthode qui doit être exécutée de manière asynchrone.
		Lorsque le runtime Go rencontre le mot-clé go, il crée une nouvelle goroutine et l'utilise pour exécuter la fonction ou la méthode spécifiée.
	    Cela modifie l'exécution du programme car, à tout moment, il existe plusieurs goroutines, chacune exécutant son propre ensemble d'instructions.
	    Ces instructions sont exécutées simultanément, ce qui signifie simplement qu'elles sont exécutées en même temps.

		Renvoi des résultats des goroutines
		Go fournit des canaux, qui sont des conduits par lesquels les données peuvent être envoyées et reçues.
		Les canaux sont fortement typés, ce qui signifie qu'ils transporteront des valeurs d'un type ou d'une interface spécifiés.
		Le type d'un canal est le mot-clé chan, suivi du type que le canal portera.
		Pour envoyer une valeur dans un canel, le canal est spécifié, suivi d'une flèche de direction exprimée par les caractères < et - puis par la valeur.
		La flèche est placée devant le canal pour en recevoir une valeur.
		Lorsqu'une valeur est envoyée via un canal, l'expéditeur n'a pas besoin de savoir comment la valeur sera reçue et utilisée.
		**/
}
