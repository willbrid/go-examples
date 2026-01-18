package main

import (
	"fmt"
	"log"
)

/**
Le package log fournit une API de journalisation simple qui crée des entrées de log et les envoie à un io.Writer, permettant à une application de générer
des données de journalisation sans avoir besoin de savoir où ces données seront stockées.
Les fonctions les plus utiles définies par le package log :
- Output() : cette fonction renvoie le Writer auquel les messages de log seront passés. Par défaut, les messages de log sont écrits sur la sortie standard.
- SetOutput(writer) : cette fonction utilise le Writer spécifié pour la journalisation.
- Flags() : cette fonction renvoie les flags utilisés pour formater les messages de journalisation.
- SetFlags(flags) : cette fonction utilise les indicateurs spécifiés pour formater les messages de journalisation.
- Prefix() : cette fonction renvoie le préfixe appliqué aux messages de journalisation. Il n'y a pas de préfixe par défaut.
- SetPrefix(prefix) : cette fonction utilise la chaîne spécifiée comme préfixe pour les messages de journalisation.
- Output(depth, message) : cette fonction écrit le message spécifié dans le Writer renvoyé par la fonction Output, avec la profondeur d'appel spécifiée,
                           qui est par défaut de 2. La profondeur d'appel est utilisée pour contrôler la sélection du fichier de code et
						   n'est généralement pas modifiée .
- Print(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint et en transmettant le résultat à la fonction Output.
- Printf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf et en transmettant le résultat à la fonction Output.
- Fatal(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint, transmet le résultat à la fonction Output, puis termine l'application.
- Fatalf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf, transmet le résultat à la fonction Output,
                              puis termine l'application.
- Panic(...vals) : cette fonction crée un message de journal en appelant fmt.Sprint, puis transmet le résultat à la fonction Output, puis à la fonction panic.
- Panicf(template, ...vals) : cette fonction crée un message de journal en appelant fmt.Sprintf et transmet le résultat à la fonction Output
                              puis à la fonction panic.

Le format des messages de log est contrôlé avec la fonction SetFlags, pour laquelle le package log définit les constantes :
- Ldate : la sélection de cet indicateur inclut la date dans la sortie du journal.
- Ltime : la sélection de cet indicateur inclut l'heure dans la sortie du journal.
- Lmicrosecondes : la sélection de cet indicateur inclut les microsecondes dans le temps.
- Llongfile : la sélection de cet indicateur inclut le nom du fichier de code, y compris les répertoires, et le numéro de ligne qui a consigné le message.
- Lshortfile : la sélection de cet indicateur inclut le nom du fichier de code, à l'exclusion des répertoires, et le numéro de ligne qui a consigné le message.
- LUTC : la sélection de cet indicateur utilise UTC pour les dates et les heures, au lieu du fuseau horaire local.
- Lmsgprefix : la sélection de cet indicateur déplace le préfixe de sa position par défaut, qui est au début du message de journal,
               juste avant la chaîne transmise à la fonction Output.
- LstdFlags : cette constante représente le format par défaut, qui sélectionne Ldate et Ltime.

Le package log peut être utilisé pour configurer différentes options de journalisation afin que différentes parties de l'application puissent écrire
des messages de log vers différentes destinations ou utiliser différentes options de formatage. La fonction ci-après est utilisée pour créer
une destination de journalisation personnalisée :
- New(writer, prefix, flags) : cette fonction renvoie un Logger qui écrira des messages au writer spécifié, configuré avec le préfixe et les flags spécifiés.
**/

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := SortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)

	log.Print("Sorted Data :", sorted)
	log.Print("Total :", total)
}

/**
La fonction d'initialisation utilise la fonction SetFlags pour sélectionner les indicateurs `Lshortfile` et `Ltime`, qui incluront
le nom du fichier et l'heure dans la sortie de journalisation. Dans la fonction main, les messages de log sont créés à l'aide de
la fonction `log.Print`.
**/

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
